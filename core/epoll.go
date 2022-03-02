//go:build linux
// +build linux

package core

import (
	"log"
	"sync"
	"syscall"

	"golang.org/x/sys/unix"
)

type Epoll struct {
	fd          int
	connections map[int]*Connection
	lock        *sync.RWMutex
}

func MkEpoll() (*Epoll, error) {
	fd, err := unix.EpollCreate1(0)
	if err != nil {
		return nil, err
	}
	return &Epoll{
		fd:          fd,
		lock:        &sync.RWMutex{},
		connections: make(map[int]*Connection),
	}, nil
}

func (e *Epoll) StartReader() {
	//读取epoll conn消息并处理
	for {
		connections, err := e.Wait()
		if err != nil {
			log.Printf("failed to Epoll wait %v", err)
			continue
		}
		for _, conn := range connections {
			if conn.Conn == nil {
				conn.Stop()
			}

			msg, err := conn.Read()
			if err != nil {
				if err := e.Remove(conn); err != nil {
					log.Printf("failed to remove %v", err)
				}
				conn.Stop()
			}
			req := Request{
				conn: *conn,
				msg:  *msg,
			}
			if msg != nil {
				conn.TCPServer.
			}

			//handle msg
		}
	}
}

func (e *Epoll) Add(conn *Connection) error {
	// Extract file descriptor associated with the connection
	fd := conn.GetConnFD()
	err := unix.EpollCtl(e.fd, syscall.EPOLL_CTL_ADD, fd, &unix.EpollEvent{Events: unix.POLLIN | unix.POLLHUP, Fd: int32(fd)})
	if err != nil {
		return err
	}
	e.lock.Lock()
	defer e.lock.Unlock()
	e.connections[fd] = conn
	log.Printf("add new conn \n")
	return nil
}

func (e *Epoll) Remove(conn *Connection) error {
	fd := conn.GetConnFD()
	err := unix.EpollCtl(e.fd, syscall.EPOLL_CTL_DEL, fd, nil)
	if err != nil {
		return err
	}
	e.lock.Lock()
	defer e.lock.Unlock()
	delete(e.connections, fd)
	if len(e.connections)%100 == 0 {
		log.Printf("total number of connections: %v", len(e.connections))
	}
	log.Printf("remove new conn \n")
	return nil
}

func (e *Epoll) Wait() ([]*Connection, error) {
	events := make([]unix.EpollEvent, 100)
retry:
	n, err := unix.EpollWait(e.fd, events, 100)
	if err != nil {
		if err == unix.EINTR {
			goto retry
		}
		return nil, err
	}
	e.lock.RLock()
	defer e.lock.RUnlock()
	var connections []*Connection
	for i := 0; i < n; i++ {
		conn := e.connections[int(events[i].Fd)]
		connections = append(connections, conn)
	}
	return connections, nil
}
