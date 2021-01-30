package server

import (
	"context"
	"log"
	"net"
	"time"
)

const (
	MaxPkg = 1024
)

type UDPServer struct {
	addr    *net.UDPAddr
	conn    *net.UDPConn
	handler Handler
	cancel  context.CancelFunc
	timeout time.Duration

	msgChan chan *UDPMsg
}

type UDPMsg struct {
	data       []byte
	remoteAddr *net.UDPAddr
	conn       *net.UDPConn
	timeout    time.Duration
	handler    Handler
}

func (u *UDPMsg) Handle(context.Context) error {
	return nil
}

func (s *UDPServer) ListenAndServer() error {
	var err error
	s.conn, err = net.ListenUDP("udp", s.addr)
	if err != nil {
		return err
	}
	defer s.conn.Close()

	ctx, cancel := context.WithCancel(context.Background())
	s.cancel = cancel
	s.msgChan = make(chan *UDPMsg)

	go s.handle(ctx)
	s.recvMsg(ctx)

	return nil
}

func (s *UDPServer) recvMsg(ctx context.Context) {
	recvBuf := make([]byte, MaxPkg)
	for {
		n, remoteAddr, err := s.conn.ReadFromUDP(recvBuf)
		if err != nil {
			log.Printf("error during read: %s", err)
		}

		msg := &UDPMsg{
			data:       make([]byte, n, n),
			remoteAddr: remoteAddr,
			conn:       s.conn,
			timeout:    s.timeout,
			handler:    s.handler,
		}
		copy(msg.data, recvBuf[:n])

		select {
		case <-ctx.Done():
			return
		case s.msgChan <- msg:
		}
	}
}

func (s *UDPServer) handle(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case m := <-s.msgChan:
			go func() {
				subCtx, cancel := context.WithTimeout(ctx, s.timeout)
				defer cancel()
				m.Handle(subCtx)
			}()
		}
	}
}
