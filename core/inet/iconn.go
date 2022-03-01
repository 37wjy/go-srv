package inet

import (
	"context"
	"net"
)

type IConn interface {
	Context() context.Context

	GetTCPConnection() *net.TCPConn
	GetConnID() string
	GetGroup() int
	GetType() string
	RemoteAddr() net.Addr

	SendMsg(msgID uint32, data []byte) error
	SendBuffMsg(msgID uint32, data []byte) error
}
