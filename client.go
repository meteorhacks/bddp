package bddp

import (
	"errors"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/glycerine/go-capnproto"
	"github.com/satori/go.uuid"
)

const (
	// Reconnection params
	ReconnInterval = time.Second * 10
	ReconnAttempts = 10

	// Ping message interval
	PingInterval = 10
)

var (
	ErrClientClosed   = errors.New("client is closed")
	ErrReconnectError = errors.New("failed to reconnect to server")
	ErrSessionFailed  = errors.New("failed to start session")
	ErrInvalidPongID  = errors.New("pong id doesn't match")
)

// Use an interface so if it's required to mock a bddp client
// for tests it can be done easily.
type Client interface {
	// Returns a channel from where the user can get errors.
	// Make sure the error handler doesn't block the goroutine.
	Errors() (errCh chan error)

	// Connects to the server if possible and also starts
	// a goroutine to reconnect if the connection drops
	Connect() (err error)

	// Closes the tcp connection to the server
	// also stops in-flight method calls
	Close() (err error)

	// Creates a MCall instnce to call a remote method
	Method(name string) (call MCall, err error)
}

type client struct {
	address string
	session string

	returnErr bool
	errChan   chan error

	conn   net.Conn
	connCh chan error

	mutex  sync.Mutex
	calls  map[string]chan *ResultMsg
	closed bool

	pingID string

	logger *log.Logger
}

func NewClient(address string) (c Client) {
	return &client{
		address: address,
		errChan: make(chan error),
		connCh:  make(chan error),
		calls:   make(map[string]chan *ResultMsg),
		logger:  log.New(os.Stderr, LogPrefix, LogFlags),
	}
}

func (c *client) Errors() (errCh chan error) {
	c.returnErr = true
	return c.errChan
}

func (c *client) Connect() (err error) {
	c.closed = true

	c.logger.Println("connecting...")

	for i := 0; i < ReconnAttempts; i++ {
		c.conn, err = net.Dial("tcp", c.address)
		if err == nil {
			c.logger.Println("connection established")
			break
		}

		// wait some time before reconnect
		c.handleErr(err)
		time.Sleep(ReconnInterval)
		c.logger.Println("re-connecting...")
	}

	if err != nil {
		return err
	}

	c.closed = false

	// automatically reconnect when
	// connection to server drops
	go c.start()

	err = c.startSession()
	if err != nil {
		c.handleErr(err)
	}

	return nil
}

func (c *client) Close() (err error) {
	if c.closed {
		return nil
	}

	c.closed = true
	c.endMethodCalls()

	return nil
}

func (c *client) Method(name string) (call MCall, err error) {
	if c.closed {
		return nil, ErrClientClosed
	}

	seg := capn.NewBuffer(nil)
	root := NewRootMessage(seg)

	call = &mcall{
		id:      uuid.NewV4().String(),
		name:    name,
		client:  c,
		segment: seg,
		message: &root,
	}

	return call, nil
}

func (c *client) start() {
	var err error

	for !c.closed {
		var msg *Message

		// stop the session if we get
		// any connection related errors
		if msg, err = c.read(); err != nil {
			break
		}

		c.processMsg(msg)
	}

	// `err == nil` only if closed by user.
	// do not reconnect when closed by user.
	if err == nil && c.closed {
		err = c.conn.Close()
		if err != nil {
			c.handleErr(err)
		}

		return
	}

	c.handleErr(err)

	c.closed = true
	// end method calls made during reconnect
	// TODO: try to recover these instead
	c.endMethodCalls()

	err = c.Connect()
	if err != nil {
		c.handleErr(ErrReconnectError)
	}
}

// end all in-flight method calls when closing client
// otherwise it might leave some goroutines hanging
// TODO: try to recover and retry instead of dropping
func (c *client) endMethodCalls() {
	c.logger.Printf("dropping %d method calls\n", len(c.calls))
	for _, ch := range c.calls {
		ch <- nil
	}
}

func (c *client) startSession() (err error) {
	seg := capn.NewBuffer(nil)
	root := NewRootMessage(seg)
	msg := NewConnectMsg(seg)
	msg.SetVersion(Version)
	support := seg.NewTextList(1)
	support.Set(0, Version)
	msg.SetSupport(support)
	root.SetConnect(msg)

	err = c.write(&root)
	if err != nil {
		return err
	}

	err = <-c.connCh
	return err
}

// TODO: use ping to check connection health
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
		err := c.write(&root)
		if err != nil {
			c.handleErr(err)
		}

		c.pingID = id
	}
}

func (c *client) read() (msg *Message, err error) {
	if c.closed {
		return nil, ErrClientClosed
	}

	seg, err := capn.ReadFromStream(c.conn, nil)
	if err != nil {
		return nil, err
	}

	root := ReadRootMessage(seg)
	return &root, nil
}

// Just write message data (cap'n proto) to the connection.
// Use a mutex to make sure messages are written one by one.
func (c *client) write(msg *Message) (err error) {
	if c.closed {
		return ErrClientClosed
	}

	seg := msg.Segment

	c.mutex.Lock()
	_, err = seg.WriteTo(c.conn)
	c.mutex.Unlock()

	return err
}

// If the session id is an empty string (no session ID)
// only accept MESSAGE_CONNECT messages. After it's set
// accept other supported message types.
func (c *client) processMsg(msg *Message) (err error) {
	mtype := msg.Which()

	switch mtype {
	case MESSAGE_CONNECTED:
		go c.handleConnected(msg)
	case MESSAGE_FAILED:
		go c.handleFailed(msg)
	case MESSAGE_PONG:
		go c.handlePong(msg)
	case MESSAGE_RESULT:
		go c.handleResult(msg)
	case MESSAGE_UPDATED:
		go c.handleUpdated(msg)
	default:
		// unknown type or corrupt msg
		c.handleErr(ErrInvalidMessage)
	}

	return nil
}

func (c *client) handleConnected(msg *Message) {
	if c.closed {
		return
	}

	req := msg.Connected()
	c.session = req.Session()
	c.connCh <- nil
}

func (c *client) handleFailed(msg *Message) {
	if c.closed {
		return
	}

	req := msg.Failed()
	ver := req.Version()
	c.logger.Println("server only supports version", ver)

	c.connCh <- ErrSessionFailed
}

func (c *client) handlePong(msg *Message) {
	if c.closed {
		return
	}

	req := msg.Pong()
	if c.pingID != req.Id() {
		c.handleErr(ErrInvalidPongID)
	}
}

func (c *client) handleResult(msg *Message) {
	if c.closed {
		return
	}

	req := msg.Result()
	id := req.Id()

	ch, ok := c.calls[id]
	if !ok {
		c.handleErr(ErrInvalidMessage)
		return
	}

	ch <- &req
}

func (c *client) handleUpdated(msg *Message) {
	if c.closed {
		return
	}

	// TODO:
	// req := msg.Updated()
}

// If user receives errors with the channel, send errors
// to the user. Otherwise log them to console STDERR.
func (c *client) handleErr(err error) {
	if c.returnErr {
		c.errChan <- err
	} else {
		c.logger.Println(err)
	}
}
