package server

import (
	"context"
)

type P2PServer struct {
}

func (s *P2PServer) Handle(context.Context, []byte) ([]byte, error) {
	return nil, nil
}
