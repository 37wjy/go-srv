package test

import (
	"UnicornServer/core/pb"
	"encoding/json"
	"testing"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

//基础struct json测试
type msgTest struct {
	Msg1 int         `json:"msg1"`
	Msg2 string      `json:"msg2"`
	Msg3 map[int]int `json:"msg3"`
}

func TestMsg(t *testing.T) {
	msg := msgTest{
		Msg1: 10,
		Msg2: "aaa",
		Msg3: map[int]int{1: 1, 2: 2},
	}
	t.Log(msg)
	data, _ := json.Marshal(msg)
	t.Logf("data: %v\n", string(data))
}

type Message = protoreflect.ProtoMessage

/*
	10001: pb.BroadCast.,
	10002: pb.Echo{},
*/
var ProtoMap = map[int]Message{
	10001: &pb.GameServerList{},
	10002: &pb.Echo{SName: "test pb"},
	10003: &pb.Error{},
}

func TestProtoMap(t *testing.T) {
	//基础反射获取
	i := ProtoMap[10002].(*pb.Echo)
	i.SHost = "123"
	t.Log(i)
	t.Log(ProtoMap[10002].ProtoReflect().Descriptor().FullName())
	t.Log(ProtoMap[10002].ProtoReflect().Descriptor())
}

func TestProtoDecode(t *testing.T) {
	// 统一decode基础
	raw_msg, _ := proto.Marshal(&pb.Echo{SHost: "dasdas", SName: "bdbdfsd", SOnline: 100})
	pb_struct := ProtoMap[10002].ProtoReflect().Type().New().Interface()
	proto.Unmarshal(raw_msg, pb_struct)
	data, _ := json.Marshal(pb_struct)
	t.Log("Unmarshal proto message :", string(data))
	data, _ = json.Marshal(ProtoMap[10002].(*pb.Echo))
	t.Log("raw proto message :", string(data))
}

/*
	测试用handler
	只是个 demo 真正的handle结构不是这样
*/

type iHandler interface {
	Handle(msgid int, data protoreflect.ProtoMessage, t *testing.T) string
}

type HandlerEcho struct{}

func (h *HandlerEcho) Handle(msgid int, data protoreflect.ProtoMessage, t *testing.T) string {
	t.Log("handler Echo input data ", data)
	msg := data.(*pb.Echo) //TODO 看看转类型还有调整空间不 范型啥的
	msg.SName = "handled"
	ret, _ := json.Marshal(msg)
	return string(ret)
}

type HandlerGameServerList struct{}

func (h *HandlerGameServerList) Handle(msgid int, data protoreflect.ProtoMessage, t *testing.T) string {
	t.Log("handler GSList input data ", data)
	msg := data.(*pb.GameServerList)
	msg.Status = 20
	ret, _ := json.Marshal(msg)
	return string(ret)
}

type HandlerError struct{}

func (h *HandlerError) Handle(msgid int, data protoreflect.ProtoMessage, t *testing.T) string {
	t.Log("handler Error input data ", data)
	msg := data.(*pb.Error)
	msg.Error = "error !!!!"
	ret, _ := json.Marshal(msg)
	return string(ret)
}

var handlerMap = map[int]iHandler{
	10001: &HandlerGameServerList{},
	10002: &HandlerEcho{},
	10003: &HandlerError{},
}

var testMsg = map[int]Message{
	10001: &pb.GameServerList{GameServerList: map[string]*pb.Server{"1232321": {SHost: "dasdas", SName: "bdbdfsd"}}},
	10002: &pb.Echo{SHost: "dasdas", SName: "bdbdfsd", SOnline: 100},
	10003: &pb.Error{Error: "no error"},
}

func doHandler(msgid int, t *testing.T) {
	//先默认没错
	proto_data, _ := testMsg[msgid]
	handle, _ := handlerMap[msgid]
	bts, _ := json.Marshal(proto_data)
	t.Log("on receive message ", msgid, ":", string(bts)) // on receive msg.....
	ret := handle.Handle(msgid, proto_data, t)
	t.Log("after process ", ret)
}

//正式handlertest
func TestHandlerMap(t *testing.T) {
	for msg := 10001; msg < 10004; msg++ {
		doHandler(msg, t)
	}
}
