package main

import (
	"github.com/DeaglePC/P2P-Server/network"
	"github.com/DeaglePC/P2P-Server/server"
)

func main() {
	p2ps := server.NewP2PServer()
	if err := network.ListenAndServerUDP("", p2ps); err != nil {
		panic(err)
	}
}
