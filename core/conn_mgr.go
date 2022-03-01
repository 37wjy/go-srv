package core

import (
	"UnicornServer/core/inet"
	"sync"
)

type ConnMgr struct {
	Lock  sync.RWMutex
	Games map[string]inet.IConn
	Room  map[string]inet.IConn
	Rank  map[string]inet.IConn
	Other map[string]inet.IConn
}

func NewConnMgr() *ConnMgr {
	return &ConnMgr{
		Games: make(map[string]inet.IConn),
		Room:  make(map[string]inet.IConn),
		Rank:  make(map[string]inet.IConn),
		Other: make(map[string]inet.IConn),
	}
}

func (s *ConnMgr) Add(conn inet.IConn) {
	switch conn.GetType() {
	case "game":
		break
	case "room":
		break
	case "rank":
		break
	default:
		s.Other[conn.GetConnID()] = conn
	}
}

func (s *ConnMgr) Remove(conn inet.IConn) {
	switch conn.GetType() {
	case "game":
		break
	case "room":
		break
	case "rank":
		break
	default:
		break
	}
}
