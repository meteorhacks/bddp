package bddp

import (
	"errors"
	"io"
	"log"
	"net"
	"time"

	"github.com/glycerine/go-capnproto"
	"github.com/satori/go.uuid"
)

const (
	ServerVersion    = "1"
	ServerBufferSize = 32
)

var (
	ErrMethodNotFound = errors.New("method not found")
)

type Server interface {
	Listen(addr string) (err error)
	Method(name string, handler MethodHandler)
	Close() (err error)
}

type server struct {
	listener net.Listener
	methods  map[string]MethodHandler
	closed   bool
}

type session struct {
	conn       net.Conn
	closed     bool
	incoming   chan *Message
	outgoing   chan *Message
	sessionId  string
	latestPing time.Time
}

type MethodContext interface {
	Segment() (seg *capn.Segment)
	Params() (params *capn.Object)
	SendResult(obj *capn.Object) (err error)
	SendError(obj *Error) (err error)
	SendUpdated() (err error)
}

type methodContext struct {
	message *Message
	result  *ResultMsg
	server  *server
	session *session
	segment *capn.Segment
	params  *capn.Object
}

type MethodHandler func(ctx MethodContext)

func NewServer() (s Server) {
	return &server{
		methods: make(map[string]MethodHandler),
	}
}

func (s *server) Listen(addr string) (err error) {
	s.listener, err = net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	for {
		if s.closed {
			break
		}

		conn, err := s.listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		ses := &session{
			conn:     conn,
			incoming: make(chan *Message, ServerBufferSize),
			outgoing: make(chan *Message, ServerBufferSize),
		}

		go s.sendMessages(ses)
		go s.recvMessages(ses)
		go s.handleConn(ses)
	}

	return nil
}

func (s *server) Method(name string, handler MethodHandler) {
	s.methods[name] = handler
}

func (s *server) Close() (err error) {
	err = s.listener.Close()
	if err != nil {
		return err
	}

	s.closed = true
	return nil
}

func (s *server) closeSession(ses *session) (err error) {
	ses.closed = true
	close(ses.incoming)
	close(ses.outgoing)

	err = ses.conn.Close()
	return err
}

func (s *server) sendMessages(ses *session) {
	for {
		root, open := <-ses.outgoing
		if !open {
			break
		}

		seg := root.Segment

		if _, err := seg.WriteTo(ses.conn); err != nil {
			log.Println(err)
		}
	}
}

func (s *server) recvMessages(ses *session) {
	for {
		root, open := <-ses.incoming
		if !open {
			break
		}

		msgType := root.Which()

		if ses.sessionId == "" {
			switch msgType {
			case MESSAGE_CONNECT:
				req := root.Connect()
				s.handleConnect(ses, &req)
			default:
				log.Println(ErrClientNotReady, msgType)
			}

			continue
		}

		switch msgType {
		case MESSAGE_PING:
			req := root.Ping()
			s.handlePing(ses, &req)
		case MESSAGE_METHOD:
			req := root.Method()
			s.handleMethod(ses, &req)
		default:
			log.Println(ErrInvalidMsgType, msgType)
		}
	}
}

func (s *server) handleConn(ses *session) {
	for {
		if ses.closed || s.closed {
			break
		}

		seg, err := capn.ReadFromStream(ses.conn, nil)
		if err == io.EOF {
			s.closeSession(ses)
			break
		} else if err != nil {
			log.Println(err)
			break
		}

		root := ReadRootMessage(seg)
		ses.incoming <- &root
	}
}

// TODO: implement resuming existing session
// TODO: implement support for multiple versions
func (s *server) handleConnect(ses *session, req *ConnectMsg) {
	seg := capn.NewBuffer(nil)
	root := NewRootMessage(seg)

	if req.Version() == ServerVersion {
		msg := NewConnectedMsg(seg)
		ses.sessionId = uuid.NewV4().String()
		msg.SetSession(ses.sessionId)
		root.SetConnected(msg)
	} else {
		msg := NewFailedMsg(seg)
		msg.SetVersion(ServerVersion)
		root.SetFailed(msg)
	}

	ses.outgoing <- &root
}

func (s *server) handlePing(ses *session, req *PingMsg) {
	seg := capn.NewBuffer(nil)
	root := NewRootMessage(seg)
	msg := NewPongMsg(seg)
	msg.SetId(req.Id())
	root.SetPong(msg)

	ses.outgoing <- &root
	ses.latestPing = time.Now()
}

func (s *server) handleMethod(ses *session, req *MethodMsg) {
	seg := capn.NewBuffer(nil)
	root := NewRootMessage(seg)
	msg := NewResultMsg(seg)
	root.SetResult(msg)
	msg.SetId(req.Id())

	name := req.Method()
	handler, ok := s.methods[name]
	if !ok {
		log.Println(ErrMethodNotFound)
		err := NewError(seg)
		msg.SetError(err)
		ses.outgoing <- &root
		return
	}

	params := req.Params()
	ctx := &methodContext{
		message: &root,
		result:  &msg,
		server:  s,
		session: ses,
		segment: seg,
		params:  &params,
	}

	handler(ctx)
}

func (c *methodContext) Segment() (segment *capn.Segment) {
	return c.segment
}

func (c *methodContext) Params() (params *capn.Object) {
	return c.params
}

func (c *methodContext) SendResult(res *capn.Object) (err error) {
	c.result.SetResult(*res)
	c.session.outgoing <- c.message

	// TODO: get and return error
	return nil
}

func (c *methodContext) SendError(obj *Error) (err error) {
	c.result.SetError(*obj)
	c.session.outgoing <- c.message

	// TODO: get and return error
	return nil
}

func (c *methodContext) SendUpdated() (err error) {
	seg := capn.NewBuffer(nil)
	root := NewRootMessage(seg)
	msg := NewUpdatedMsg(seg)
	root.SetUpdated(msg)
	c.session.outgoing <- &root

	// TODO: get and return error
	return nil
}
