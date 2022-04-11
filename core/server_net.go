package core

import (
	"UnicornServer/core/logger"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	_ "net/http/pprof"
)

type ServerCfg struct {
	Name   string
	IP     string
	Port   int
	LogDir string
}

func getCfg() *ServerCfg {
	config := &ServerCfg{}
	f, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		logger.Info("use default config")
		return &ServerCfg{
			Name: "UnicornCenter",
			IP:   "0.0.0.0",
			Port: 9998}
	}
	err = json.Unmarshal(f, config)
	if err != nil {
		logger.Fatal("unmarshal faild ", err)
	}
	return config
}

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
	cfg := getCfg()
	s := &Server{
		Name:      cfg.Name,
		IPVersion: "tcp4",
		IP:        cfg.IP,
		Port:      cfg.Port,
		ConnMgr:   NewConnMgr(),
		Handler:   NewMsgHandle(),
	}
	return s
}

func (s *Server) Start() {
	logger.Infof("%s start serving at %s:%d \n", s.Name, s.IP, s.Port)

	go func() {

		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			logger.Fatal("resolve tcp addr err: ", err)
			return
		}

		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			panic(err)
		}

		//TODO server.go 应该有一个自动生成ID的方法

		//3 启动server网络连接业务
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				logger.Fatal("Accept connection err ", err)
				continue
			}
			logger.Info("Accept connection remote addr = ", conn.RemoteAddr().String())

			dealConn := NewConnection(s, conn)

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
