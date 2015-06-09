package bddp

import (
	"github.com/glycerine/go-capnproto"
	"github.com/satori/go.uuid"
)

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
