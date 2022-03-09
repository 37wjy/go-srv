package core

type Request struct {
	conn *Connection
	msg  *Message
}

func (r *Request) GetConnection() *Connection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgID()
}

func (r *Request) SetMsgID(id uint32) {
	r.msg.SetMsgID(id)
}
