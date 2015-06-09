package server

import (
	"errors"
	"io"
	"net"
	"sync"

	"github.com/glycerine/go-capnproto"
	"github.com/meteorhacks/bddp"
	"github.com/satori/go.uuid"
)

var (
	ErrSessionClosed  = errors.New("client's session is closed")
	ErrConnected      = errors.New("client is already connected")
	ErrNotConnected   = errors.New("client is not connected")
	ErrInvalidMessage = errors.New("invalid message type")
	ErrMethodNotFound = errors.New("method not found")
)

type session struct {
	id     string
	conn   net.Conn
	mutex  sync.Mutex
	server *server
	closed bool
}

func newSession(conn net.Conn, server *server) (s *session) {
	return &session{
		conn:   conn,
		server: server,
	}
}

// Closes the session and prevents future reads and writes
func (s *session) close() (err error) {
	if s.closed {
		return nil
	}

	s.closed = true
	return nil
}

// Read a bddp message from the tcp connection
func (s *session) read() (msg *bddp.Message, err error) {
	if s.closed {
		return nil, ErrSessionClosed
	}

	seg, err := capn.ReadFromStream(s.conn, nil)
	if err != nil {
		return nil, err
	}

	root := bddp.ReadRootMessage(seg)
	return &root, nil
}

// Just write message data (cap'n proto) to the connection.
// Use a mutex to make sure messages are written one by one.
func (s *session) write(msg *bddp.Message) (err error) {
	if s.closed {
		return ErrSessionClosed
	}

	seg := msg.Segment

	s.mutex.Lock()
	_, err = seg.WriteTo(s.conn)
	s.mutex.Unlock()

	return err
}

func (s *session) process() (err error) {
	for !s.closed {
		var msg *bddp.Message

		// stop the session if we get
		// any connection related errors
		if msg, err = s.read(); err != nil {
			break
		}

		s.processMsg(msg)
	}

	// EOF usually means a disconnect
	if err != io.EOF {
		return err
	}

	return nil
}

// If the session id is an empty string (no session ID)
// only accept MESSAGE_CONNECT messages. After it's set
// accept other supported message types.
func (s *session) processMsg(msg *bddp.Message) (err error) {
	mtype := msg.Which()

	switch mtype {
	case bddp.MESSAGE_CONNECT:
		s.handleConnect(msg)
	case bddp.MESSAGE_PING:
		go s.handlePing(msg)
	case bddp.MESSAGE_METHOD:
		go s.handleMethod(msg)
	default:
		// unknown type or corrupt msg
		s.handleErr(ErrInvalidMessage)
	}

	return nil
}

// TODO: implement resuming existing session
// TODO: implement support for multiple versions
func (s *session) handleConnect(msg *bddp.Message) {
	req := msg.Connect()

	seg := capn.NewBuffer(nil)
	root := bddp.NewRootMessage(seg)

	if req.Version() == ServerVersion {
		res := bddp.NewConnectedMsg(seg)
		s.id = uuid.NewV4().String()
		res.SetSession(s.id)
		root.SetConnected(res)
	} else {
		res := bddp.NewFailedMsg(seg)
		res.SetVersion(ServerVersion)
		root.SetFailed(res)
	}

	s.write(&root)
}

func (s *session) handlePing(msg *bddp.Message) {
	if s.closed {
		return
	}

	req := msg.Ping()

	seg := capn.NewBuffer(nil)
	root := bddp.NewRootMessage(seg)
	res := bddp.NewPongMsg(seg)
	res.SetId(req.Id())
	root.SetPong(res)

	s.write(&root)
}

func (s *session) handleMethod(msg *bddp.Message) {
	if s.closed {
		return
	}

	req := msg.Method()

	seg := capn.NewBuffer(nil)
	root := bddp.NewRootMessage(seg)
	res := bddp.NewResultMsg(seg)
	root.SetResult(res)
	res.SetId(req.Id())

	name := req.Method()
	handler, ok := s.server.methods[name]

	// send an error if the method
	// handler does not exist
	if !ok {
		s.handleErr(ErrMethodNotFound)
		err := bddp.NewError(seg)
		res.SetError(err)
		s.write(&root)
		return
	}

	params := req.Params()
	ctx := &mcontext{
		method:  name,
		message: &root,
		result:  &res,
		session: s,
		segment: seg,
		params:  &params,
	}

	handler(ctx)
}

// Forward all errors to the servers handleErr
// From there, user can receive them if needed
func (s *session) handleErr(err error) {
	s.server.handleErr(err)
}
