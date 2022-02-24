package main

import (
	"flag"
	"log"
	_ "net/http/pprof"
	"os"
	"time"

	"UnicornServer/core"

	"github.com/rcrowley/go-metrics"
)


func main() {

	go metrics.Log(metrics.DefaultRegistry, 5*time.Second, log.New(os.Stderr, "metrics: ", log.Lmicroseconds))

	s = core.NewServer()
}
