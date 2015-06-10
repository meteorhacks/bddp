package bddp

import (
	"errors"

	"github.com/glycerine/go-capnproto"
)

var (
	ErrMethodInterrupted = errors.New("client disconnected")
)

type MCall interface {
	Segment() (seg *capn.Segment)
	Call(params capn.Object) (res capn.Object, err error)
}

type mcall struct {
	id      string
	name    string
	client  *client
	segment *capn.Segment
	message *Message
}

func (m *mcall) Segment() (seg *capn.Segment) {
	return m.segment
}

// Response will be nil if the method call fails inflight
func (m *mcall) Call(params capn.Object) (res capn.Object, err error) {
	root := m.message
	msg := NewMethodMsg(m.segment)
	msg.SetId(m.id)
	msg.SetMethod(m.name)
	msg.SetParams(params)
	root.SetMethod(msg)

	ch := make(chan *ResultMsg)
	m.client.calls[m.id] = ch
	m.client.write(root)

	// wait until we get a response
	response := <-ch
	delete(m.client.calls, m.id)

	if response == nil {
		err = ErrMethodInterrupted
		return res, err
	}

	switch response.Which() {
	case RESULTMSG_RESULT:
		res = response.Result()
	case RESULTMSG_ERROR:
		err = response.Error()
	}

	return res, err
}
