package main

import (
	"zinx/net"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
}

func main() {
	server := net.NewServer("test")
	server.Serve()
}
