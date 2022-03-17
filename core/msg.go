package core

type Message struct {
	ID      int32
	DataLen uint32 //消息的长度
	Data    []byte //消息的内容
}

//NewMsgPackage 创建一个Message消息包
func NewMsgPackage(ID int32, data []byte) *Message {
	return &Message{
		ID:      ID,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}

func (msg *Message) GetMsgID() int32 {
	return msg.ID
}

func (msg *Message) SetMsgID(id int32) {
	msg.ID = id
}

//GetDataLen 获取消息数据段长度
func (msg *Message) GetDataLen() uint32 {
	return msg.DataLen
}

//GetData 获取消息内容
func (msg *Message) GetData() []byte {
	return msg.Data
}

//SetDataLen 设置消息数据段长度
func (msg *Message) SetDataLen(len uint32) {
	msg.DataLen = len
}

//SetData 设计消息内容
func (msg *Message) SetData(data []byte) {
	msg.Data = data
}
