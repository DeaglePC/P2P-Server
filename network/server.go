package network

import "context"

type Server interface {
	ListenAndServer() error
}

type Handler interface {
	Handle(context.Context, []byte) ([]byte, error)
}
