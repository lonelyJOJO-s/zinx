package net

import (
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"
)

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
}

func NewServer(name string) *Server {
	return &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "127.0.0.1",
		Port:      8999,
	}
}

func (s *Server) Start() {
	// 开启服务
	// 1. 解析套接字
	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		log.Println("resovle addr error:", err)
		return
	}
	log.Println("resolve addr success")
	// 2. 建立listener
	lisener, err := net.ListenTCP(s.IPVersion, addr)
	if err != nil {
		log.Println("listen connection failed:", err)
		return
	}
	log.Println("lisener have been build successfully")
	// 3. 监听连接
	for {
		// 如果connect
		coon, err := lisener.AcceptTCP()
		if err != nil {
			log.Println("coon failed:", err)
			continue
		}

		// question: for?
		go func() {
			for {
				buf := make([]byte, 512)
				cnt, err := coon.Read(buf)
				if err != nil {
					log.Println("coon read error:", err)
					return
				}
				log.Println("read info:", string(buf[:cnt]))
				if _, err := coon.Write(buf[:cnt]); err != nil {
					log.Println("write back error:", err)
					return
				}
			}
		}()
	}
}

func (s *Server) Stop() {
	// 关停服务
}

func (s *Server) Serve() {
	// 服务
	go s.Start()
	// todo: additional business
	select {}
}
