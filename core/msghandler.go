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

	req.conn.SendMsg(123, req.msg.Data)
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
	msg := &pb.BroadCast{}
	err := proto.Unmarshal(req.GetData(), msg)
	if err != nil {
		logger.Fatal("handshake Unmarshal error ", err)
		return
	}
}

func process_echo(req *Request) {

}

func process_transfer(req *Request) {

}
