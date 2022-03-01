package core

import (
	"context"
	"fmt"
	"io"
	"net"
	"sync"
)

type Connection struct {
	TCPServer  Server
	Epoller    Epoll
	Conn       *net.TCPConn
	ConnGroup  uint32
	ConnID     uint32
	ConnType   string
	ConnBranch string

	ctx     context.Context
	cancel  context.CancelFunc
	msgChan chan []byte

	sync.RWMutex
	isClosed bool
}

func NewConnection(server Server, epoll Epoll, conn *net.TCPConn) *Connection {
	c := &Connection{
		TCPServer: server,
		Epoller:   epoll,
		Conn:      conn,
		ConnID:    1,
		ConnType:  "game",

		msgChan:  make(chan []byte),
		isClosed: false,
	}
	//c.TCPServer.
	return c
}

func (c *Connection) StartWriter() {
	fmt.Println("[Writer Goroutine is running]")
	defer fmt.Println(c.RemoteAddr().String(), "[conn Writer exit!]")

	for {
		select {
		case data := <-c.msgChan:
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("Send Data error:, ", err, " Conn Writer exit")
				return
			}
		case <-c.ctx.Done():
			return
		}
	}
}

func (c *Connection) StartReader() {
	fmt.Println("[Reader Goroutine is running!]")
	defer c.Stop() //读爆了就停
	defer fmt.Println(c.RemoteAddr().String(), "[conn Reader exit]")
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			headData := make([]byte, getHeaderLen())
			if _, err := io.ReadFull(c.Conn, headData); err != nil {
				fmt.Println("read msg head err", err)
				return
			}

			/*
				msg, error := decode(headData)
				if err != nil {
					fmt.Println()
				}
			*/
			////改 TODO
			msg, err := c.TCPServer.Packet().Unpack(headData)
			if err != nil {
				fmt.Println("unpack error ", err)
				return
			}

			//根据 dataLen 读取 data，放在msg.Data中
			var data []byte
			if msg.GetDataLen() > 0 {
				data = make([]byte, msg.GetDataLen())
				if _, err := io.ReadFull(c.Conn, data); err != nil {
					fmt.Println("read msg data error ", err)
					return
				}
			}
			msg.SetData(data)

			//得到当前客户端请求的Request数据
			req := Request{
				conn: c,
				msg:  msg,
			}

			if utils.GlobalObject.WorkerPoolSize > 0 {
				//已经启动工作池机制，将消息交给Worker处理
				c.MsgHandler.SendMsgToTaskQueue(&req)
			} else {
				//从绑定好的消息和对应的处理方法中执行对应的Handle方法
				go c.MsgHandler.DoMsgHandler(&req)
			}
		}
	}
}

func (c *Connection) Start() {
	c.StartReader()

	c.StartWriter()

	select {
	case <-c.ctx.Done():
		c.Finalizer()
		return
	}
}

func (c *Connection) Stop() {
	c.cancel()
}

func (c *Connection) SendMsg(data []byte) error {
	return nil
}

func (c *Connection) Finalizer() {
	c.Lock()
	defer c.Unlock()

	_ = c.Conn.Close()

	c.isClosed = true
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) GetGroup() uint32 {
	return c.ConnGroup
}

func (c *Connection) GetType() string {
	return c.ConnType
}
