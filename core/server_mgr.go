package core


import (
	"UnicornServer/core/inet"
)




type ServerMgr struct {
	Games 	map[string]inet.IConnection
	Games 	map[string]inet.IConnection
	Games 	map[string]inet.IConnection
}

//NewServerNgr 创建ServerNgr
func NewServerNgr() *ServerMgr {
	return &ServerMgr{
		Apis:           make(map[uint32]inet.IRouter),
		WorkerPoolSize: inet.GlobalObject.WorkerPoolSize,
		//一个worker对应一个queue
		TaskQueue: make([]chan ziface.IRequest, utils.GlobalObject.WorkerPoolSize),
	}
}

