package core

import (
	"context"
	"fmt"
	"io"
	"net"
	"sync"
)

type Connection struct {
	TCPServer  *Server
	Conn       net.Conn
	ConnGroup  uint32
	ConnType   string
	ConnBranch string

	ctx     context.Context
	cancel  context.CancelFunc
	msgChan chan []byte

	sync.RWMutex
	isClosed bool
}

func NewConnection(server *Server, conn net.Conn) *Connection {
	c := &Connection{
		TCPServer: server,
		Conn:      conn,
		ConnType:  "game",

		msgChan:  make(chan []byte),
		isClosed: false,
	}
	//c.TCPServer.
	return c
}

func (c *Connection) StartReader() {
	fmt.Println("[Reader Goroutine is running]")
	defer fmt.Println(c.RemoteAddr().String(), "[conn Reader exit!]")
	defer c.Stop()

	// 创建拆包解包的对象
	for {
		select {
		case <-c.ctx.Done():
			return
		default:

			//读取客户端的Msg head
			headData := make([]byte, getHeaderLen())
			if _, err := io.ReadFull(c.Conn, headData); err != nil {
				fmt.Println("read msg head error ", err)
				return
			}

			msg, err := Unpack(headData)
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

			go c.TCPServer.Handler.DoMsgHandler(&req)

		}
	}
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

func (c *Connection) Start() {
	c.StartWriter()
	c.StartReader()

	select {
	case <-c.ctx.Done():
		c.Finalizer()
		return
	}
}

func (c *Connection) Stop() {
	c.cancel()
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

func (c *Connection) GetGroup() uint32 {
	return c.ConnGroup
}

func (c *Connection) GetType() string {
	return c.ConnType
}
