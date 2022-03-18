package core

import (
	"UnicornServer/core/logger"
	"UnicornServer/core/pb"

	"google.golang.org/protobuf/proto"
)

// MsgHandle -
type MsgHandle struct {
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{}
}

func (mh *MsgHandle) DoMsgHandler(req *Request) {

	preprocess(req)
	defer postprocess(req)

	//req.conn.SendMsg(123, req.msg.Data)
	//handler 量大时需实现继承ihandler
	switch req.GetMsgID() {
	case 10001:
		process_handshake(req)
		return
	case 10002:
		process_echo(req)
		return
	default:
		process_transfer(req)
		return
	}
}

//handlers
//这么写吧 不想写interface了
func preprocess(req *Request) {
	logger.Infof("process msg %d", req.GetMsgID())
}

func postprocess(req *Request) {

}

func process_handshake(req *Request) {
	msg := &pb.HandShake{}
	err := proto.Unmarshal(req.GetData(), msg)
	if err != nil {
		logger.Fatal("handshake Unmarshal error ", err)
		return
	}

	conn := req.conn
	conn.ConnName = msg.SName
	conn.ConnHost = msg.SHost
	conn.ConnGroup = msg.SGroup
	conn.ConnBranch = msg.CurrBranch
	logger.Infof("receive connection from %s, %s, group: %d", msg.SName, msg.SHost, msg.SGroup)
	req.conn.TCPServer.ConnMgr.Add(conn)

	switch msg.SName {
	case "game":
		ret_sp, _ := proto.Marshal(&pb.SpecialServerList{
			RoomServerList: conn.TCPServer.ConnMgr.GetRoomServer(),
			RankServerList: conn.TCPServer.ConnMgr.GetRankServer(),
		})
		conn.SendMsg(MsgID.SPECIAL_SERVER_LIST, ret_sp) //special server list
		ret_gs, _ := proto.Marshal(&pb.GameServerList{Status: 1, GameServerList: map[string]*pb.Server{msg.SHost: {SHost: msg.SHost, SGroup: &msg.SGroup}}})
		conn.TCPServer.ConnMgr.SyncGameServer(MsgID.GAME_SERVER_LIST, ret_gs)
	case "room":
		//sync game server
		ret_gs, _ := proto.Marshal(&pb.GameServerList{Status: 0, GameServerList: conn.TCPServer.ConnMgr.GetGameServer()})
		conn.TCPServer.ConnMgr.SyncGameServer(MsgID.GAME_SERVER_LIST, ret_gs)
		//sync room server
		conn.TCPServer.ConnMgr.SyncRoomServer()
	case "rank":
		//sync game server
		ret_gs, _ := proto.Marshal(&pb.GameServerList{Status: 0, GameServerList: conn.TCPServer.ConnMgr.GetGameServer()})
		conn.TCPServer.ConnMgr.SyncGameServer(MsgID.GAME_SERVER_LIST, ret_gs)
		//sync rank server
		conn.TCPServer.ConnMgr.SyncRankServer()
	default:
		//add gm
	}
}

func process_echo(req *Request) {
	msg := &pb.Echo{}
	err := proto.Unmarshal(req.GetData(), msg)
	if err != nil {
		logger.Fatal("echo Unmarshal error ", err)
		return
	}
	logger.Info("processing echo")
	conn := req.conn
	if conn.GetName() == "game" {
		conn.Online = *msg.SOnline
	}
}

func process_transfer(req *Request) {
	msg := &pb.BroadCast{}
	err := proto.Unmarshal(req.GetData(), msg)
	if err != nil {
		logger.Fatal("trans Unmarshal error ", err)
		return
	}
	logger.Infof("processing transfer %d from %s to %s", req.GetMsgID(), req.conn.GetName(), msg.Target)
	req.conn.TCPServer.ConnMgr.SendToHost(msg.Target, req.GetMsgID(), req.GetData())
}
