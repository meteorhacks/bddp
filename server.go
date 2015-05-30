package bddp

import (
	"errors"
	"io"
	"log"
	"net"
	"time"

	"github.com/glycerine/go-capnproto"
)

var (
	ErrUnknownMsgType = errors.New("unknown message type")
)

type Server interface {
	Listen(addr string) (err error)
	Close() (err error)
}

type server struct {
	listener net.Listener
	closed   bool
}

type session struct {
	conn       net.Conn
	latestPing time.Time
}

func NewServer() (s Server) {
	return &server{}
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

		go s.handleConn(conn)
	}

	return nil
}

func (s *server) Close() (err error) {
	err = s.listener.Close()
	if err != nil {
		return err
	}

	s.closed = true
	return nil
}

func (s *server) handleConn(conn net.Conn) {
	ses := &session{conn: conn}

	for {
		if s.closed {
			break
		}

		seg, err := capn.ReadFromStream(conn, nil)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println(err)
			break
		}

		root := ReadRootMessage(seg)

		switch t := root.Which(); t {
		case MESSAGE_PING:
			req := root.Ping()
			s.handlePing(ses, &req)
		default:
			log.Println(ErrUnknownMsgType, t)
		}
	}
}

func (s *server) handlePing(ses *session, req *PingMsg) {
	seg := capn.NewBuffer(nil)
	root := NewRootMessage(seg)
	msg := NewPongMsg(seg)
	msg.SetId(req.Id())
	root.SetPong(msg)

	if _, err := seg.WriteTo(ses.conn); err != nil {
		log.Println(err)
		return
	}

	ses.latestPing = time.Now()
}
