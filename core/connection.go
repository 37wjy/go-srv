package core

import (
	"context"
	"fmt"
	"io"
	"net"
	"sync"
)

type Connection struct {
	TCPServer Server
	Conn      *net.TCPConn
	ConnID    uint32

	ctx     context.Context
	cancel  context.CancelFunc
	msgChan chan []byte

	sync.RWMutex
	property     map[string]interface{}
	propertyLock sync.Mutex
	isClosed     bool
}

func NewConnection(server Server, conn *net.TCPConn) *Connection {
	c := &Connection{
		TCPServer: server,
		Conn:      conn,

		msgChan:  make(chan []byte),
		isClosed: false,
		property: nil,
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
	defer c.Stop()
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

			msg,error:=decode(headData)
			if err!=nil{
				fmt.Println()
			}
		}
	}
}

func (c *Connection) Start() {

}

func (c *Connection) Stop() {

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
