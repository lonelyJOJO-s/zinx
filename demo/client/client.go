package main

import (
	"net"
	"time"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
}

func main() {

	coon, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		log.Info("connect to tcp failed, error:", err)
		return
	}

	for {
		_, err := coon.Write([]byte("hello guys"))
		if err != nil {
			log.Info("write error:", err)
			return
		}
		buf := make([]byte, 512)
		cnt, err := coon.Read(buf)
		if err != nil {
			log.Info("read error:", err)
			return
		}
		log.Info(string(buf[:cnt]))
		time.Sleep(time.Second)
	}

}
