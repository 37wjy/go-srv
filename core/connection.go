package core

import (
	"context"
	"fmt"
	"io"
	"net"
	"reflect"
	"sync"
)

type Connection struct {
	TCPServer  *Server
	Epoller    *Epoll
	Conn       net.Conn
	ConnGroup  uint32
	fd         int
	ConnType   string
	ConnBranch string

	ctx     context.Context
	cancel  context.CancelFunc
	msgChan chan []byte

	sync.RWMutex
	isClosed bool
}

func socketFD(conn net.Conn) int {
	tcpConn := reflect.Indirect(reflect.ValueOf(conn)).FieldByName("conn")
	fdVal := tcpConn.FieldByName("fd")
	pfdVal := reflect.Indirect(fdVal).FieldByName("pfd")

	return int(pfdVal.FieldByName("Sysfd").Int())
}

func NewConnection(server *Server, epoll *Epoll, conn net.Conn) *Connection {
	c := &Connection{
		TCPServer: server,
		Epoller:   epoll,
		Conn:      conn,
		fd:        socketFD(conn),
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

func (c *Connection) Read() (*Message, error) {
	headData := make([]byte, getHeaderLen())
	if _, err := io.ReadFull(c.Conn, headData); err != nil {
		fmt.Println("read msg head err", err)
		return nil, err
	}

	msg, err := Unpack(headData)
	if err != nil {
		fmt.Println("unpack error ", err)
		return nil, err
	}

	//根据 dataLen 读取 data，放在msg.Data中
	var data []byte
	if msg.GetDataLen() > 0 {
		data = make([]byte, msg.GetDataLen())
		if _, err := io.ReadFull(c.Conn, data); err != nil {
			fmt.Println("read msg data error ", err)
			return nil, err
		}
	}
	msg.SetData(data)

	return &msg, nil
}

func (c *Connection) Start() {
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
	c.Epoller.Remove(c)
	_ = c.Conn.Close()

	c.isClosed = true
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) GetConnFD() int {
	return c.fd
}

func (c *Connection) GetGroup() uint32 {
	return c.ConnGroup
}

func (c *Connection) GetType() string {
	return c.ConnType
}
