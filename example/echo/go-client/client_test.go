package main

import (
	"testing"
)

func BenchmarkEcho(b *testing.B) {
	str := "echo"
	c := Prepare()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		SendEcho(c, str)
	}
}
