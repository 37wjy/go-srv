package inet

import (
	"context"
	"net"
)

type IServer interface {
	Start()
	Stop()
	Context() context.Context

	GetTCPConnection() *net.TCPConn
	GetConnID() string
	RemoteAddr() net.Addr

	SendMsg(msgID uint32, data []byte) error
	SendBuffMsg(msgID uint32, data []byte) error
}
