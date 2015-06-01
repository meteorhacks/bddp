package main

import (
	"github.com/meteorhacks/bddp"
)

func main() {
	Prepare()
	select {}
}

func Prepare() (s bddp.Server) {
	// create a server and listen
	s = bddp.NewServer()
	go s.Listen(":3300")

	s.Method("echo", func(ctx bddp.MethodContext) {
		params := ctx.Params()
		str := params.ToText()
		obj := ctx.Segment().NewText(str)
		ctx.SendResult(&obj)
		ctx.SendUpdated()
	})

	return s
}
