package core

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"google.golang.org/protobuf/proto"
)

var (
	HeadLen uint16 = 4
)

func getHeaderLen() uint16 {
	return HeadLen
}

func getLen() {

}

func Unpack(binaryData []byte) (*Message, error) {
	dataBuff := bytes.NewReader(binaryData)

	msg := &Message{}

	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	return msg, nil
}

func Pack(msg *Message) ([]byte, error) {
	dataBuff := bytes.NewBuffer([]byte{})

	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}

	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil
}

func aa() {
	//1. 将客户端传来的proto协议解码
	msg := &pb.Position{}
	err := proto.Unmarshal(request.GetData(), msg)
	if err != nil {
		fmt.Println("Move: Position Unmarshal error ", err)
		return
	}

	//2. 得知当前的消息是从哪个玩家传递来的,从连接属性pID中获取
	pID, err := request.GetConnection().GetProperty("pID")
	if err != nil {
		fmt.Println("GetProperty pID error", err)
		request.GetConnection().Stop()
		return
	}

	//fmt.Printf("user pID = %d , move(%f,%f,%f,%f)\n", pID, msg.X, msg.Y, msg.Z, msg.V)

	//3. 根据pID得到player对象
	player := core.WorldMgrObj.GetPlayerByPID(pID.(int32))

	//4. 让player对象发起移动位置信息广播
	player.UpdatePos(msg.X, msg.Y, msg.Z, msg.V)
}
