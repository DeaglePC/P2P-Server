package client

import "net"

type Info struct {
	ID   uint64
	Name string

	LastHeartbeatTime int64

	UDPAddr *net.UDPAddr
}

type Client interface {
	Store(info *Info) error
	Load(id uint64) (*Info, error)
	Delete(id uint64) error
	UpdateHeartbeat(id uint64) error
}
