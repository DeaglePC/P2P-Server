package main

import (
	"flag"
	"os"

	"github.com/DeaglePC/P2P-Server/network"
	"github.com/DeaglePC/P2P-Server/server"
	log "github.com/sirupsen/logrus"
)

var (
	addr = flag.String("addr", ":10086", "addr")
)

func init() {
	flag.Parse()

	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
}

func main() {
	p2ps := server.NewP2PServer()
	if err := network.ListenAndServerUDP(*addr, p2ps); err != nil {
		panic(err)
	}
}
