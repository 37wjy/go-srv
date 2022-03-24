package core

import (
	"UnicornServer/core/logger"
	"UnicornServer/core/pb"
	"sync"

	"google.golang.org/protobuf/proto"
)

type ConnMgr struct {
	Lock  sync.RWMutex
	games map[string]*Connection
	rooms map[string]*Connection
	ranks map[string]*Connection
	other map[string]*Connection
	hosts map[string]*Connection
}

func NewConnMgr() *ConnMgr {
	return &ConnMgr{
		games: make(map[string]*Connection),
		rooms: make(map[string]*Connection),
		ranks: make(map[string]*Connection),
		other: make(map[string]*Connection),
		hosts: make(map[string]*Connection),
	}
}

func (c *ConnMgr) Add(conn *Connection) {
	switch conn.GetName() {
	case "game":
		c.games[conn.GetHost()] = conn
	case "room":
		c.rooms[conn.GetHost()] = conn
	case "rank":
		c.ranks[conn.GetHost()] = conn
	default:
		c.other[conn.GetHost()] = conn
	}
	c.hosts[conn.GetHost()] = conn
}

func (c *ConnMgr) Remove(conn *Connection) {
	switch conn.GetName() {
	case "game":
		delete(c.games, conn.GetHost())
	case "room":
		delete(c.rooms, conn.GetHost())
	case "rank":
		delete(c.ranks, conn.GetHost())
	default:
		delete(c.other, conn.GetHost())
	}
	delete(c.hosts, conn.GetHost())
}

func (c *ConnMgr) SendToHost(host string, msgID int32, data []byte) {
	conn, ok := c.hosts[host]
	if !ok {
		logger.Fatal("send msg to host faild: ", host)
		return
	}
	conn.SendMsg(msgID, data)
}

func (c *ConnMgr) SendToAllGame(msgID int32, data []byte) {
	for _, conn := range c.games {
		conn.SendMsg(msgID, data)
	}
}

func (c *ConnMgr) SendToAllRoom(msgID int32, data []byte) {
	for _, conn := range c.rooms {
		conn.SendMsg(msgID, data)
	}
}

func (c *ConnMgr) SendToAllRank(msgID int32, data []byte) {
	for _, conn := range c.ranks {
		conn.SendMsg(msgID, data)
	}
}

func (c *ConnMgr) SyncGameServer(msgID int32, data []byte) {
	c.SendToAllRank(msgID, data)
	c.SendToAllRoom(msgID, data)
}

func (c *ConnMgr) SyncRoomServer() {
	data, _ := proto.Marshal(
		&pb.SpecialServerList{
			RankServerList: nil,
			RoomServerList: c.GetRoomServer(),
		},
	)
	c.SendToAllGame(MsgID.SPECIAL_SERVER_LIST, data)
}

func (c *ConnMgr) SyncRankServer() {
	data, _ := proto.Marshal(
		&pb.SpecialServerList{
			RankServerList: c.GetRankServer(),
			RoomServerList: nil,
		},
	)
	c.SendToAllGame(MsgID.SPECIAL_SERVER_LIST, data)
}

func (c *ConnMgr) GetGameServer() map[string]*pb.Server {
	ret := make(map[string]*pb.Server)
	for k, conn := range c.games {
		ret[k] = &pb.Server{SHost: conn.GetHost(), SGroup: &conn.ConnGroup}
	}
	return ret
}

func (c *ConnMgr) GetRoomServer() map[string]*pb.Server {
	ret := make(map[string]*pb.Server)
	for k, conn := range c.rooms {
		ret[k] = &pb.Server{SHost: conn.GetHost()}
	}
	return ret
}

func (c *ConnMgr) GetRankServer() map[string]*pb.Server {
	ret := make(map[string]*pb.Server)
	for k, conn := range c.ranks {
		ret[k] = &pb.Server{SHost: conn.GetHost()}
	}
	return ret
}

func (c *ConnMgr) GMGetGameServer() map[string]*pb.Server {
	ret := make(map[string]*pb.Server)
	for k, conn := range c.games {
		ret[k] = &pb.Server{SHost: conn.GetHost(), SGroup: &conn.ConnGroup}
	}
	return ret
}

func (c *ConnMgr) GMGetRoomServer() map[string]*pb.Server {
	ret := make(map[string]*pb.Server)
	for k, conn := range c.rooms {
		ret[k] = &pb.Server{SHost: conn.GetHost()}
	}
	return ret
}

func (c *ConnMgr) GMGetRankServer() map[string]*pb.Server {
	ret := make(map[string]*pb.Server)
	for k, conn := range c.ranks {
		ret[k] = &pb.Server{SHost: conn.GetHost()}
	}
	return ret
}
