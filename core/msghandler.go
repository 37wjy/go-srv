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
	case 10005:
		process_gm_get_setver_list(req)
		return
	case 10006:
		process_test(req)
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
	conn.ConnName = msg.GetSName()
	conn.ConnHost = msg.GetSHost()
	conn.ConnGroup = msg.GetSGroup()
	conn.ConnBranch = msg.GetCurrBranch()
	logger.Infof("receive connection from %s, %s, group: %d", msg.GetSName(), msg.GetSHost(), msg.GetSGroup())
	req.conn.TCPServer.ConnMgr.Add(conn)
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
		conn.Online = msg.GetSOnline()
	}
	ret, _ := proto.Marshal(&pb.Echo{})
	conn.SendMsg(MsgID.ECHO, ret)
}

func process_gm_get_setver_list(req *Request) {
	logger.Info("processing gm get server list")
	conn := req.conn

	ret, _ := proto.Marshal(&pb.GMServerList{
		GameServerList: conn.TCPServer.ConnMgr.GMGetGameServer(),
		RoomServerList: conn.TCPServer.ConnMgr.GMGetRoomServer(),
		RankServerList: conn.TCPServer.ConnMgr.GMGetRankServer(),
	})
	conn.SendMsg(MsgID.GAME_SERVER_LIST, ret)

}

func process_transfer(req *Request) {
	msg := &pb.BroadCast{}
	err := proto.Unmarshal(req.GetData(), msg)
	if err != nil {
		logger.Fatal("trans Unmarshal error ", err)
		return
	}
	//现役GM用
	if req.GetMsgID() >= MsgID.GM_ID_START && req.GetMsgID() < MsgID.RK_ID_START {
		if req.conn.GetName() != "gm" {
			logger.Infof("processing transfer %d from %s to GM", req.GetMsgID(), req.conn.GetName())
			req.conn.TCPServer.ConnMgr.SendToAllOthers(req.GetMsgID(), req.GetData())
			return
		}
	}
	//转发
	switch msg.GetOpertype() {
	case pb.BroadCast_BROADCAST_GAME:
		logger.Infof("processing broadcast %d from %s to game", req.GetMsgID(), req.conn.GetName())
		req.conn.TCPServer.ConnMgr.SendToAllGame(req.GetMsgID(), req.GetData())
	case pb.BroadCast_BROADCAST_ROOM:
		logger.Infof("processing broadcast %d from %s to room", req.GetMsgID(), req.conn.GetName())
		req.conn.TCPServer.ConnMgr.SendToAllRoom(req.GetMsgID(), req.GetData())
	case pb.BroadCast_BROADCAST_RANK:
		logger.Infof("processing broadcast %d from %s to rank", req.GetMsgID(), req.conn.GetName())
		req.conn.TCPServer.ConnMgr.SendToAllRank(req.GetMsgID(), req.GetData())
	default:
		logger.Infof("processing transfer %d from %s to %s", req.GetMsgID(), req.conn.GetName(), msg.GetTarget())
		req.conn.TCPServer.ConnMgr.SendToHost(msg.GetTarget(), req.GetMsgID(), req.GetData())
	}

}

func process_test(req *Request) {
	msg := &pb.Test{}
	conn := req.conn
	err := proto.Unmarshal(req.GetData(), msg)
	logger.Infof("on receive msg %+v", msg)
	if err != nil {
		logger.Fatal("trans Unmarshal error ", err)
		return
	}
	rlist := msg.GetIlist()
	rmap := msg.GetStrmap()
	if rlist == nil {
		rlist = make([]int32, 0)
	}
	if rmap == nil {
		rmap = make(map[string]string)
	}
	if msg.GetMsgp() != 0 {
		rlist = append(rlist, msg.GetMsgp())
	}
	if msg.GetMsgk() != "" && msg.GetMsgv() != "" {
		rmap[msg.GetMsgk()] = msg.GetMsgv()
	}

	ret, _ := proto.Marshal(&pb.Test{
		Strmap: rmap,
		Ilist:  rlist,
	})
	conn.SendMsg(10006, ret)
}
