package main

import (
	"errors"
	"time"

	"github.com/meteorhacks/bddp"
)

var (
	ErrEchoDoesntMatch = errors.New("strings doesn't match")
)

func main() {
	c := Prepare()
	for t := range time.Tick(time.Second) {
		str := t.String()
		SendEcho(c, str)
	}
}

func Prepare() (c bddp.Client) {
	// create a client and connect
	c = bddp.NewClient()
	if err := c.Connect("localhost:3300"); err != nil {
		panic(err)
	}

	return c
}

func SendEcho(c bddp.Client, str string) (err error) {
	call := c.NewMethodCall("echo")
	seg := call.Segment()
	params := seg.NewText("test-echo-string")

	res, err := call.Call(params)
	if err != nil {
		return err
	}

	out := res.ToText()
	if out != str {
		return ErrEchoDoesntMatch
	}

	return nil
}
