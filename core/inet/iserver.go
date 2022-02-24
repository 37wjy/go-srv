package inet

import (
	"context"
	"net"
)

type IServer interface {
	Context() context.Context

	GetTCPConnection() *net.TCPConn
	GetConnID() string
	GetGroup() string
	RemoteAddr() net.Addr

	SendMsg(msgID uint32, data []byte) error
	SendBuffMsg(msgID uint32, data []byte) error
}
