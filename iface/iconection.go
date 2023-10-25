package iface

import "net"

type IConnection interface {
	// 启动链接
	Start()
	// 结束链接
	Stop()
	// 获取当前连接的sockect
	GetTCPConnection() *net.TCPConn
	// 获取当前id
	GetCoonID() uint32
	// 获取对端状态
	RemoteAddr() net.Addr
	// 发送数据
	Send(data []byte) error
}

type HandleFunc func(*net.TCPConn, []byte, int) error
