package net

import (
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
}

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
		IP:        "0.0.0.0",
		Port:      8999,
	}
}

func (s *Server) Start() {
	go func() {
		// 1. acquire tcp addr
		log.Printf("[Start] Server listen at IP: %s, Port: %d is starting\n", s.IP, s.Port)
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			log.Info("reslove tcp addr error: ", err)
			return
		}

		// 2. listen to the server
		listner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			log.Info("listen ", s.IPVersion, "err: ", err)
			return
		}
		log.Info("start server succ, ", s.Name, "succ, Listening...")

		// 3. wait for connecting
		for {
			conn, err := listner.AcceptTCP()
			if err != nil {
				log.Info("Accept err: ", err)
				continue
			}
			// conn has been established
			go func(conn *net.TCPConn) {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						log.Info("recv buf error: ", err)
						continue
					}
					if _, err := conn.Write(buf[:cnt]); err != nil {
						log.Info("wirte back error: ", err)
						continue
					}

				}
			}(conn)
		}
	}()
}

func (s *Server) Stop() {

}

func (s *Server) Serve() {
	s.Start()
	// block
	// TODO: 做一些额外业务
	select {}
}
