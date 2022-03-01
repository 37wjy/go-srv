package core

import (
	"fmt"
	"log"
	"net"
	_ "net/http/pprof"
	"syscall"

	"github.com/libp2p/go-reuseport"
	"github.com/rcrowley/go-metrics"
)

var (
	c       = 1
	opsRate = metrics.NewRegisteredMeter("ops", nil)
)

type Server struct {
	Name        string
	IPVersion   string
	IP          string
	Port        int
	epolls      []Epoll
	ConnMgr     *ConnMgr
	OnConnStart func(conn net.Conn)
	OnConnStop  func(conn net.Conn)
}

func NewServer() *Server {
	setLimit()
	s := &Server{
		Name:      "UnicornCenter",
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      9999,
		ConnMgr:   NewConnMgr(),
	}
	return s
}

func (s *Server) Start() {
	fmt.Println("%s start serving at %s:%d", s.Name, s.IP, s.Port)
	for i := 0; i < c; i++ {
		go s.startEpoll()
	}
	select {}
}

func (s *Server) startEpoll() {
	ln, err := reuseport.Listen("tcp", ":9999")
	if err != nil {
		panic(err)
	}

	epoller, err := MkEpoll()
	if err != nil {
		panic(err)
	}

	go epoller.Start()

	for {
		conn, e := ln.Accept()

		if e != nil {
			if ne, ok := e.(net.Error); ok && ne.Temporary() {
				log.Printf("accept temp err: %v", ne)
				continue
			}

			log.Printf("accept err: %v", e)
			return
		}

		if err := epoller.Add(conn); err != nil {
			log.Printf("failed to add connection %v", err)
			conn.Close()
		}
	}
}

func (s *Server) SetOnConnectionAdd(hookFunc func(conn net.Conn)) {
	s.OnConnStart = hookFunc
}

func (s *Server) SetOnConnectionLost(hookFunc func(conn net.Conn)) {
	s.OnConnStop = hookFunc
}

func setLimit() {
	var rLimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}
	rLimit.Cur = rLimit.Max
	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}

	log.Printf("set cur limit: %d", rLimit.Cur)
}
