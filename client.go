package bddp

import (
	"errors"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/glycerine/go-capnproto"
	"github.com/satori/go.uuid"
)

const (
	PingInterval     = 10
	ClientVersion    = "1"
	ClientBufferSize = 32
)

var (
	ErrClientClosed   = errors.New("client is closed")
	ErrClientNotReady = errors.New("client not connected")
	ErrInvalidPongId  = errors.New("ping-pong ids doesn't match")
	ErrInvalidMsgType = errors.New("invalid message type")
	ErrConnectFailed  = errors.New("failed to create session")
	ErrInvalidResult  = errors.New("invalid method result")
)

type Client interface {
	Connect(addr string) (err error)
	NewMethodCall(name string) (call MethodCall)
}

type client struct {
	conn       net.Conn
	closed     bool
	incoming   chan *Message
	outgoing   chan *Message
	calls      map[string]chan *ResultMsg
	connectCh  chan error
	pingId     string
	sessionId  string
	latestPong time.Time
}

type MethodCall interface {
	Segment() (seg *capn.Segment)
	Call(params capn.Object) (res capn.Object, err error)
}

type methodCall struct {
	id      string
	name    string
	client  *client
	segment *capn.Segment
	message *Message
}

func NewClient() (c Client) {
	return &client{
		incoming:  make(chan *Message, ClientBufferSize),
		outgoing:  make(chan *Message, ClientBufferSize),
		calls:     make(map[string]chan *ResultMsg),
		connectCh: make(chan error),
	}
}

func (c *client) Connect(addr string) (err error) {
	c.conn, err = net.Dial("tcp", addr)
	if err != nil {
		return err
	}

	go c.sendMessages()
	go c.recvMessages()
	go c.handleConn()
	go c.startSession()
	go c.startPing()

	// TODO: timeout
	err = <-c.connectCh
	return err
}

func (c *client) Close() (err error) {
	close(c.incoming)
	close(c.outgoing)
	c.closed = true

	err = c.conn.Close()
	return err
}

func (c *client) sendMessages() {
	for {
		root, open := <-c.outgoing
		if !open {
			break
		}

		seg := root.Segment

		if _, err := seg.WriteTo(c.conn); err != nil {
			log.Println(err)
		}
	}
}

func (c *client) recvMessages() {
	for {
		root, open := <-c.incoming
		if !open {
			break
		}

		msgType := root.Which()

		if c.sessionId == "" {
			switch msgType {
			case MESSAGE_CONNECTED:
				req := root.Connected()
				c.handleConnected(&req)
			case MESSAGE_FAILED:
				req := root.Failed()
				c.handleFailed(&req)
			default:
				log.Println(ErrClientNotReady, msgType)
			}

			continue
		}

		switch msgType {
		case MESSAGE_PONG:
			req := root.Pong()
			c.handlePong(&req)
		case MESSAGE_RESULT:
			req := root.Result()
			c.handleResult(&req)
		case MESSAGE_UPDATED:
			req := root.Updated()
			c.handleUpdated(&req)
		default:
			log.Println(ErrInvalidMsgType, msgType)
		}
	}
}

func (c *client) startSession() {
	seg := capn.NewBuffer(nil)
	root := NewRootMessage(seg)
	msg := NewConnectMsg(seg)
	msg.SetVersion(ClientVersion)
	support := seg.NewTextList(1)
	support.Set(0, ClientVersion)
	msg.SetSupport(support)
	root.SetConnect(msg)

	c.outgoing <- &root
}

func (c *client) startPing() {
	counter := 0

	seg := capn.NewBuffer(nil)
	root := NewRootMessage(seg)
	msg := NewPingMsg(seg)
	root.SetPing(msg)

	for _ = range time.Tick(time.Second * PingInterval) {
		if c.closed {
			break
		}

		id := strconv.Itoa(counter)
		counter++

		msg.SetId(id)
		c.outgoing <- &root
		c.pingId = id
	}
}

func (c *client) handleConn() {
	for {
		if c.closed {
			break
		}

		seg, err := capn.ReadFromStream(c.conn, nil)
		if err == io.EOF {
			c.Close()
			break
		} else if err != nil {
			log.Println(err)
			break
		}

		root := ReadRootMessage(seg)
		c.incoming <- &root
	}
}

func (c *client) handlePong(req *PongMsg) {
	if c.pingId != req.Id() {
		log.Println(ErrInvalidPongId)
		return
	}

	c.latestPong = time.Now()
}

func (c *client) handleConnected(req *ConnectedMsg) {
	c.sessionId = req.Session()
	c.connectCh <- nil
}

func (c *client) handleFailed(req *FailedMsg) {
	c.connectCh <- ErrConnectFailed
}

func (c *client) handleResult(req *ResultMsg) {
	id := req.Id()
	ch, ok := c.calls[id]
	if !ok {
		log.Println(ErrInvalidResult)
		return
	}

	ch <- req
}

func (c *client) handleUpdated(req *UpdatedMsg) {
	// TODO
}

func (c *client) NewMethodCall(name string) (call MethodCall) {
	id := uuid.NewV4().String()
	seg := capn.NewBuffer(nil)
	root := NewRootMessage(seg)

	return &methodCall{id, name, c, seg, &root}
}

func (m *methodCall) Segment() (seg *capn.Segment) {
	return m.segment
}

func (m *methodCall) Call(params capn.Object) (res capn.Object, err error) {
	root := m.message
	msg := NewMethodMsg(m.segment)
	msg.SetId(m.id)
	msg.SetMethod(m.name)
	msg.SetParams(params)
	root.SetMethod(msg)

	ch := make(chan *ResultMsg)
	m.client.calls[m.id] = ch
	m.client.outgoing <- root

	// wait until we get a response
	response := <-ch
	delete(m.client.calls, m.id)

	switch response.Which() {
	case RESULTMSG_RESULT:
		res = response.Result()
	case RESULTMSG_ERROR:
		err = response.Error()
	}

	return res, err
}
