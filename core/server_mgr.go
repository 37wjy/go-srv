package core

import (
	"UnicornServer/core/inet"
	"net"
)

type ServerMgr struct {
	Games map[string]inet.IServer
	Room  map[string]inet.IServer
	Rank  map[string]inet.IServer
	GM    map[string]inet.IServer
}

//NewServerNgr 创建ServerNgr
func NewServerNgr() *ServerMgr {
	return &ServerMgr{
		Games: make(map[string]inet.IServer),
		Room:  make(map[string]inet.IServer),
		Rank:  make(map[string]inet.IServer),
		GM:    make(map[string]inet.IServer),
	}
}

func (s *ServerMgr) Add(conn net.Conn) {

}

func (s *ServerMgr) Remove(conn net.Conn) {

}
