package bddp

import (
	"errors"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/glycerine/go-capnproto"
)

const (
	PingSeconds = 10
)

var (
	ErrClientClosed  = errors.New("client is closed")
	ErrInvalidPongId = errors.New("ping-pong ids doesn't match")
)

type Client interface {
	Connect(addr string) (err error)
}

type client struct {
	conn       net.Conn
	closed     bool
	pingId     string
	latestPong time.Time
}

func NewClient() (c Client) {
	return &client{}
}

func (c *client) Connect(addr string) (err error) {
	c.conn, err = net.Dial("tcp", addr)
	if err != nil {
		return err
	}

	go c.handleConn()
	go c.startPing()

	return nil
}

func (c *client) Close() (err error) {
	err = c.conn.Close()
	if err != nil {
		return err
	}

	c.closed = true
	return nil
}

func (c *client) startPing() {
	counter := 0

	for _ = range time.Tick(time.Second * PingSeconds) {
		id := strconv.Itoa(counter)
		counter++

		if c.closed {
			break
		}

		seg := capn.NewBuffer(nil)
		root := NewRootMessage(seg)
		msg := NewPingMsg(seg)
		msg.SetId(id)
		root.SetPing(msg)

		if _, err := seg.WriteTo(c.conn); err != nil {
			log.Println(err)
			continue
		}

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
			break
		} else if err != nil {
			log.Println(err)
			break
		}

		root := ReadRootMessage(seg)

		switch t := root.Which(); t {
		case MESSAGE_PONG:
			req := root.Pong()
			c.handlePong(&req)
		default:
			log.Println(ErrUnknownMsgType, t)
		}
	}
}

func (c *client) handlePong(req *PongMsg) {
	if c.pingId != req.Id() {
		log.Println(ErrInvalidPongId)
		return
	}

	c.latestPong = time.Now()
}
