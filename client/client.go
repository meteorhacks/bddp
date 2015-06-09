package client

import (
	"errors"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/glycerine/go-capnproto"
	"github.com/meteorhacks/bddp"
	"github.com/satori/go.uuid"
)

const (
	// BDDP protocol version used by the client.
	// Until the client supports multiple versions,
	// client version should match server version.
	Version = "1"

	// Reconnection params
	ReconnInterval = time.Second * 5

	// Ping message interval
	PingInterval = 10

	// Error logger params
	LogPrefix = "bddp: "
	LogFlags  = log.LstdFlags
)

var (
	ErrClientClosed   = errors.New("client is closed")
	ErrInvalidMessage = errors.New("invalid message type")
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
	calls  map[string]chan *bddp.ResultMsg
	closed bool

	pingID string

	logger *log.Logger
}

func New(address string) (c Client) {
	return &client{
		address: address,
		errChan: make(chan error),
		connCh:  make(chan error),
		calls:   make(map[string]chan *bddp.ResultMsg),
		logger:  log.New(os.Stderr, LogPrefix, LogFlags),
	}
}

func (c *client) Errors() (errCh chan error) {
	c.returnErr = true
	return c.errChan
}

func (c *client) Connect() (err error) {
	c.conn, err = net.Dial("tcp", c.address)
	if err != nil {
		return err
	}

	// automatically reconnect when
	// connection to server drops
	go c.reconnect()

	err = c.startSession()
	if err != nil {
		return err
	}

	return nil
}

func (c *client) Close() (err error) {
	if c.closed {
		return nil
	}

	c.closed = true

	// end all in-flight method calls when closing client
	// otherwise it might leave some goroutines hanging
	for _, ch := range c.calls {
		ch <- nil
	}

	return nil
}

func (c *client) Method(name string) (call MCall, err error) {
	seg := capn.NewBuffer(nil)
	root := bddp.NewRootMessage(seg)

	call = &mcall{
		id:      uuid.NewV4().String(),
		name:    name,
		client:  c,
		segment: seg,
		message: &root,
	}

	return call, nil
}

func (c *client) reconnect() {
	var err error

	err = c.process()
	if err != nil {
		c.handleErr(err)
	}

	err = c.Close()
	if err != nil {
		c.handleErr(err)
	}

	// wait some time before reconnect
	time.Sleep(ReconnInterval)

	// TODO: for better memory performance, recover client
	// struct instead of replacing it with a new one
	newc := New(c.address).(*client)

	err = newc.Connect()
	if err != nil {
		newc.handleErr(err)
	}

	*c = *newc
}

func (c *client) startSession() (err error) {
	seg := capn.NewBuffer(nil)
	root := bddp.NewRootMessage(seg)
	msg := bddp.NewConnectMsg(seg)
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

func (c *client) startPing() {
	counter := 0

	seg := capn.NewBuffer(nil)
	root := bddp.NewRootMessage(seg)
	msg := bddp.NewPingMsg(seg)
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

func (c *client) read() (msg *bddp.Message, err error) {
	if c.closed {
		return nil, ErrClientClosed
	}

	seg, err := capn.ReadFromStream(c.conn, nil)
	if err != nil {
		return nil, err
	}

	root := bddp.ReadRootMessage(seg)
	return &root, nil
}

// Just write message data (cap'n proto) to the connection.
// Use a mutex to make sure messages are written one by one.
func (c *client) write(msg *bddp.Message) (err error) {
	if c.closed {
		return ErrClientClosed
	}

	seg := msg.Segment

	c.mutex.Lock()
	_, err = seg.WriteTo(c.conn)
	c.mutex.Unlock()

	return err
}

func (c *client) process() (err error) {
	for !c.closed {
		var msg *bddp.Message

		// stop the session if we get
		// any connection related errors
		if msg, err = c.read(); err != nil {
			break
		}

		c.processMsg(msg)
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
func (c *client) processMsg(msg *bddp.Message) (err error) {
	mtype := msg.Which()

	switch mtype {
	case bddp.MESSAGE_CONNECTED:
		c.handleConnected(msg)
	case bddp.MESSAGE_FAILED:
		c.handleFailed(msg)
	case bddp.MESSAGE_PONG:
		c.handlePong(msg)
	case bddp.MESSAGE_RESULT:
		c.handleResult(msg)
	case bddp.MESSAGE_UPDATED:
		c.handleUpdated(msg)
	default:
		// unknown type or corrupt msg
		c.handleErr(ErrInvalidMessage)
	}

	return nil
}

func (c *client) handleConnected(msg *bddp.Message) {
	if c.closed {
		return
	}

	req := msg.Connected()
	c.session = req.Session()
	c.connCh <- nil
}

func (c *client) handleFailed(msg *bddp.Message) {
	if c.closed {
		return
	}

	req := msg.Failed()
	ver := req.Version()
	c.logger.Println("server only supports version", ver)

	c.connCh <- ErrSessionFailed
}

func (c *client) handlePong(msg *bddp.Message) {
	if c.closed {
		return
	}

	req := msg.Pong()
	if c.pingID != req.Id() {
		c.handleErr(ErrInvalidPongID)
	}
}

func (c *client) handleResult(msg *bddp.Message) {
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

func (c *client) handleUpdated(msg *bddp.Message) {
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
