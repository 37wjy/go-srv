package main

import (
	_ "net/http/pprof"

	"UnicornServer/core"
)

func main() {

	s := core.NewServer()

	s.Start()
}
