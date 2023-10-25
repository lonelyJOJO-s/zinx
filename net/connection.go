package net

import (
	"log"
	"net"
	"zinx/iface"
)

type Connection struct {
	Coon     *net.TCPConn
	CoonID   uint32
	isClosed bool
	// 连接绑定的业务方法
	handleAPI iface.HandleFunc
	// 告知连接停止channel
	ExitChan chan bool
}

func NewConnection(coon *net.TCPConn, coonID uint32, callbackAPI iface.HandleFunc) *Connection {
	return &Connection{
		Coon:      coon,
		CoonID:    coonID,
		isClosed:  false,
		handleAPI: callbackAPI,
		ExitChan:  make(chan bool, 1),
	}
}

func (c *Connection) Start() {
	log.Println("Coon start... conn id:", c.CoonID)
	// 启动读业务

	// todo: 启动写业务
}

func (c *Connection) Stop() {
	log.Println("Coon stop... conn id:", c.CoonID)
	if c.isClosed {
		return
	}
	// 关闭连接
	c.isClosed = true
	c.Coon.Close()
	close(c.ExitChan)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Coon
}

func (c *Connection) GetCoonID() uint32 {
	return c.CoonID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Coon.RemoteAddr()
}

func (c *Connection) Send(data []byte) (err error) {
	return
}
