package main

import (
	"UnicornServer/core"
	_ "net/http/pprof"
)

func main() {
	s := core.NewServer()
	s.Start()
}
