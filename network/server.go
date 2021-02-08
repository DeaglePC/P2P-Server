package network

import (
	"context"
	"net"
)

type Server interface {
	ListenAndServer() error
}

type Context struct {
	context.Context
	RemoteAddr *net.UDPAddr
	Conn       *net.UDPConn
}

type Handler interface {
	Handle(Context, []byte) ([]byte, error)
}

func ListenAndServerUDP(addr string, handler Handler) error {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return err
	}
	server := UDPServer{
		addr:    udpAddr,
		handler: handler,
	}
	return server.ListenAndServer()
}
