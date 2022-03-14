package core

import (
	"UnicornServer/core/logger"
	"context"
	"errors"
	"io"
	"net"
	"sync"
)

type Connection struct {
	TCPServer  *Server
	Conn       net.Conn
	ConnHost   string
	ConnGroup  int32
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

		msgChan:  make(chan []byte),
		isClosed: false,
	}
	//c.TCPServer.
	return c
}

func (c *Connection) StartReader() {
	logger.Debug("[Reader Goroutine is running]")
	defer logger.Debug(c.RemoteAddr().String() + "[conn Reader exit!]")
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
				logger.Fatal("read msg head error ", err)
				return
			}

			msg, err := UnpackHead(headData)
			if err != nil {
				logger.Fatal("unpack error ", err)
				return
			}

			//根据 dataLen 读取 data，放在msg.Data中
			var data []byte
			if msg.GetDataLen() > 0 {
				data = make([]byte, msg.GetDataLen())
				if _, err := io.ReadFull(c.Conn, data); err != nil {
					logger.Fatal("read msg data error ", err)
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
	logger.Debug("[Writer Goroutine is running]")
	defer logger.Debug(c.RemoteAddr().String(), "[conn Writer exit!]")

	for {
		select {
		case data := <-c.msgChan:
			if _, err := c.Conn.Write(data); err != nil {
				logger.Fatal("Send Data error:, ", err, " Conn Writer exit")
				return
			}
		case <-c.ctx.Done():
			return
		}
	}
}

func (c *Connection) Start() {
	c.ctx, c.cancel = context.WithCancel(context.Background())
	go c.StartWriter()
	go c.StartReader()

	select {
	case <-c.ctx.Done():
		c.Finalizer()
		return
	}
}

func (c *Connection) Stop() {
	logger.Infof("Connection %s close", c.GetHost())
	c.cancel()
}

func (c *Connection) Finalizer() {
	c.Lock()
	c.TCPServer.ConnMgr.Remove(c)
	defer c.Unlock()
	_ = c.Conn.Close()

	c.isClosed = true
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) GetGroup() int32 {
	return c.ConnGroup
}

func (c *Connection) GetHost() string {
	return c.ConnHost
}

func (c *Connection) GetType() string {
	return c.ConnType
}

func (c *Connection) SendMsg(msgID uint32, data []byte) error {
	c.RLock()
	defer c.RUnlock()
	if c.isClosed == true {
		return errors.New("connection closed when send msg")
	}

	//将data封包，并且发送
	msg, err := Pack(NewMsgPackage(msgID, data))
	if err != nil {
		logger.Fatal("Pack error msg ID = ", msgID)
		return errors.New("Pack error msg ")
	}

	//写回客户端
	c.msgChan <- msg

	return nil
}
