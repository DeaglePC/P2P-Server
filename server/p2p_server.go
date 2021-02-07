package server

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/DeaglePC/P2P-Server/network"
	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
)

const ClientTimeoutSec = 5

type ClientInfo struct {
	ID   uint64
	Name string

	LastHeartbeatTime int64

	UDPAddr *net.UDPAddr
}

type P2PServer struct {
	clients *sync.Map
}

func NewP2PServer() *P2PServer {
	return &P2PServer{clients: new(sync.Map)}
}

func (s *P2PServer) deleteClient(id uint64) {
	s.clients.Delete(id)
	log.Infof("delete client %d", id)
}

func (s *P2PServer) checkHeartbeat() {
	var once sync.Once
	once.Do(func() {
		go func() {
			for {
				s.clients.Range(func(key, value interface{}) bool {
					if time.Now().Unix()-value.(*ClientInfo).LastHeartbeatTime > ClientTimeoutSec {
						s.deleteClient(key.(uint64))
					}
					return true
				})
				time.Sleep(time.Duration(1) * time.Second)
			}
		}()
	})
}

func (s *P2PServer) Handle(ctx network.Context, data []byte) ([]byte, error) {
	msg := &Msg{}
	if err := proto.Unmarshal(data, msg); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	var req, resp proto.Message
	respMsg := &Msg{Header: &MsgHeader{}}
	if err := proto.Unmarshal(msg.Body, req); err != nil {
		return nil, err
	}

	switch msg.Header.Type {
	case MsgType_HEARTBEAT_REQ:
		resp = s.updateHeartbeat(ctx, req)
		respMsg.Header.Type = MsgType_HEARTBEAT_RESP
	case MsgType_LOGIN_REQ:
		resp = s.login(ctx, req)
		respMsg.Header.Type = MsgType_LOGIN_RESP
	}

	respBytes, err := proto.Marshal(resp)
	if err != nil {
		return nil, err
	}
	respMsg.Body = respBytes

	respMsgBytes, err := proto.Marshal(respMsg)
	if err != nil {
		return nil, err
	}

	return respMsgBytes, nil
}

func (s *P2PServer) updateHeartbeat(ctx network.Context, req proto.Message) (resp proto.Message) {
	hbReq := req.(*HeartbeatReq)
	resp = &HeartbeatResp{ID: hbReq.ID}

	client, ok := s.clients.Load(hbReq.ID)
	if !ok {
		resp.(*HeartbeatResp).Result = RetCode_UserNotFound
		return
	}
	client.(*ClientInfo).LastHeartbeatTime = time.Now().Unix()

	s.checkHeartbeat()
	return
}

func (s *P2PServer) login(ctx network.Context, req proto.Message) (resp proto.Message) {
	loginReq := req.(*LoginReq)

	id := getID()
	s.clients.Store(id, ClientInfo{
		ID:                id,
		Name:              loginReq.Name,
		UDPAddr:           ctx.RemoteAddr,
		LastHeartbeatTime: time.Now().Unix(),
	})

	return &LoginResp{
		ID: id,
	}
}

func (s *P2PServer) logout(ctx network.Context, req proto.Message) (resp proto.Message) {
	logoutReq := req.(*LogoutReq)
	resp = &LogoutResp{}

	_, ok := s.clients.Load(logoutReq.ID)
	if !ok {
		resp.(*LogoutResp).Result = RetCode_UserNotFound
		return
	}
	s.deleteClient(logoutReq.ID)

	return
}

func (s *P2PServer) punch(ctx network.Context, req proto.Message) (resp proto.Message) {
	return
}
