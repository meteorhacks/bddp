package bddp

import (
	"github.com/glycerine/go-capnproto"
)

type MethodContext interface {
	Method() (name string)
	Segment() (seg *capn.Segment)
	Params() (params *capn.Object)
	SendResult(obj *capn.Object) (err error)
	SendError(obj *Error) (err error)
	SendUpdated() (err error)
}

type methodContext struct {
	method  string
	message *Message
	result  *ResultMsg
	server  *server
	session *session
	segment *capn.Segment
	params  *capn.Object
}

func (c *methodContext) Method() (name string) {
	return c.method
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
