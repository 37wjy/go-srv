package core

import (
	"fmt"
	"net"
	_ "net/http/pprof"

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
	ConnMgr     *ConnMgr
	Handler     *MsgHandle
	OnConnStart func(conn net.Conn)
	OnConnStop  func(conn net.Conn)
}

func NewServer() *Server {

	s := &Server{
		Name:      "UnicornCenter",
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      9999,
		ConnMgr:   NewConnMgr(),
		Handler:   NewMsgHandle(),
	}
	return s
}

func (s *Server) Start() {
	fmt.Println("%s start serving at %s:%d", s.Name, s.IP, s.Port)
	for i := 0; i < c; i++ {

	}
	select {}
}

func (s *Server) SetOnConnectionAdd(hookFunc func(conn net.Conn)) {
	s.OnConnStart = hookFunc
}

func (s *Server) SetOnConnectionLost(hookFunc func(conn net.Conn)) {
	s.OnConnStop = hookFunc
}
