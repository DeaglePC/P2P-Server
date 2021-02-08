package client

import "net"

type Client struct {
	ID                uint64
	Name              string
	LastHeartbeatTime int64
	UDPAddr           *net.UDPAddr
}

type Set interface {
	Store(info *Client) error
	Load(id uint64) (*Client, error)
	Delete(id uint64) error
	UpdateHeartbeat(id uint64) error
	DeleteTimeOut(timeout int64) error
}
