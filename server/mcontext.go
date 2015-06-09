package server

import (
	"github.com/glycerine/go-capnproto"
	"github.com/meteorhacks/bddp"
)

type MContext interface {
	Method() (name string)
	Segment() (seg *capn.Segment)
	Params() (params *capn.Object)
	SendResult(obj *capn.Object) (err error)
	SendError(obj *bddp.Error) (err error)
	SendUpdated() (err error)
}

type mcontext struct {
	method  string
	message *bddp.Message
	result  *bddp.ResultMsg
	session *session
	segment *capn.Segment
	params  *capn.Object
}

func (c *mcontext) Method() (name string) {
	return c.method
}

func (c *mcontext) Segment() (segment *capn.Segment) {
	return c.segment
}

func (c *mcontext) Params() (params *capn.Object) {
	return c.params
}

func (c *mcontext) SendResult(res *capn.Object) (err error) {
	c.result.SetResult(*res)
	err = c.session.write(c.message)
	return err
}

func (c *mcontext) SendError(obj *bddp.Error) (err error) {
	c.result.SetError(*obj)
	err = c.session.write(c.message)
	return err
}

func (c *mcontext) SendUpdated() (err error) {
	seg := capn.NewBuffer(nil)
	root := bddp.NewRootMessage(seg)
	msg := bddp.NewUpdatedMsg(seg)
	root.SetUpdated(msg)
	err = c.session.write(c.message)
	return err
}
