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
	fmt.Printf("%s start serving at %s:%d \n", s.Name, s.IP, s.Port)
	go func() {

		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err: ", err)
			return
		}

		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			panic(err)
		}

		//TODO server.go 应该有一个自动生成ID的方法

		//3 启动server网络连接业务
		for {
			//3.1 阻塞等待客户端建立连接请求
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err ", err)
				continue
			}
			fmt.Println("Get conn remote addr = ", conn.RemoteAddr().String())

			//3.3 处理该新连接请求的 业务 方法， 此时应该有 handler 和 conn是绑定的
			dealConn := NewConnection(s, conn)

			//3.4 启动当前链接的处理业务
			go dealConn.Start()
		}
	}()
	select {}
}

func (s *Server) SetOnConnectionAdd(hookFunc func(conn net.Conn)) {
	s.OnConnStart = hookFunc
}

func (s *Server) SetOnConnectionLost(hookFunc func(conn net.Conn)) {
	s.OnConnStop = hookFunc
}
