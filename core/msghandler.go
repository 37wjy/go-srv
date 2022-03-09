package core

import (
	"UnicornServer/core/pb"
	"fmt"

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
	req.conn.SendMsg(123, req.msg.Data)
	switch req.GetMsgID() {
	case 10001:
		return
	case 10002:
		return
	default:
		return

	}
}

//这么写吧 不想写interface了
func preprocess(req *Request) {

}

func postprocess(req *Request) {

}

func processhandshake(req *Request) {
	msg := &pb.BroadCast{}
	err := proto.Unmarshal(req.GetData(), msg)
	if err != nil {
		fmt.Println("Move: Position Unmarshal error ", err)
		return
	}
}

func processecho(req *Request) {

}

func processtransfer(req *Request) {

}
