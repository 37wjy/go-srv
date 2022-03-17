package inet

import (
	"context"
	"net"
)

type IConn interface {
	Context() context.Context

	GetConnID() string
	GetHost() string
	GetGroup() int
	GetName() string
	RemoteAddr() net.Addr

	SendMsg(msgID int32, data []byte) error
}
