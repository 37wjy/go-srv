package core

import (
	"sync"
)

type ConnMgr struct {
	Lock  sync.RWMutex
	Games map[string]*Connection
	Room  map[string]*Connection
	Rank  map[string]*Connection
	Other map[string]*Connection
}

func NewConnMgr() *ConnMgr {
	return &ConnMgr{
		Games: make(map[string]*Connection),
		Room:  make(map[string]*Connection),
		Rank:  make(map[string]*Connection),
		Other: make(map[string]*Connection),
	}
}

func (c *ConnMgr) Add(conn *Connection) {
	switch conn.GetType() {
	case "game":
		c.Games[conn.GetHost()] = conn
		break
	case "room":
		c.Room[conn.GetHost()] = conn
		break
	case "rank":
		c.Rank[conn.GetHost()] = conn
		break
	default:
		c.Other[conn.GetHost()] = conn
	}
}

func (c *ConnMgr) Remove(conn *Connection) {
	switch conn.GetType() {
	case "game":
		delete(c.Games, conn.GetHost())
		break
	case "room":
		delete(c.Room, conn.GetHost())
		break
	case "rank":
		delete(c.Rank, conn.GetHost())
		break
	default:
		delete(c.Other, conn.GetHost())
		break
	}
}

func (s *ConnMgr) SendToHost(host string, msgID uint32, data []byte) {

}

func (s *ConnMgr) SendToAllGame(msgID uint32, data []byte) {

}

func (s *ConnMgr) SendToAllRoom(msgID uint32, data []byte) {

}

func (s *ConnMgr) SendToAllRank(msgID uint32, data []byte) {

}
