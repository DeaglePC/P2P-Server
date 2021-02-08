package server

import (
	"fmt"
	"sync"
	"time"

	c "github.com/DeaglePC/P2P-Server/client"
	"github.com/DeaglePC/P2P-Server/network"
	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
)

const ClientTimeoutSec = 5

type P2PServer struct {
	clients c.Set
}

func NewP2PServer() *P2PServer {
	return &P2PServer{clients: c.NewClientWithMemStore()}
}

func (s *P2PServer) deleteClient(id uint64) {
	if err := s.clients.Delete(id); err != nil {
		log.Errorf("delete clients err: %+v", err)
	}
	log.Infof("delete clients %d", id)
}

func (s *P2PServer) checkHeartbeat() {
	var once sync.Once
	once.Do(func() {
		go func() {
			for {
				if err := s.clients.DeleteTimeOut(ClientTimeoutSec); err != nil {
					log.Error("checkHeartbeat err: %+v", err)
				}
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

	if err := s.clients.UpdateHeartbeat(hbReq.ID); err != nil {
		log.Errorf("updateHeartbeat err: %+v", err)
	}

	s.checkHeartbeat()
	return
}

func (s *P2PServer) login(ctx network.Context, req proto.Message) (resp proto.Message) {
	loginReq := req.(*LoginReq)

	id := getID()
	_ = s.clients.Store(&c.Info{
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

	if err := s.clients.Delete(logoutReq.ID); err != nil {
		resp.(*LogoutResp).Result = RetCode_UserNotFound
	}
	return
}

func (s *P2PServer) punch(ctx network.Context, req proto.Message) (resp proto.Message) {
	punchReq := req.(*PunchReq)
	resp = &PunchResp{}

	client, err := s.clients.Load(punchReq.ID)
	if err != nil {
		resp.(*PunchResp).Result = RetCode_UserNotFound
		log.Errorf("cann't find cliend: %d, err: %+v", punchReq.ID, err)
		return
	}

	targetClient, err := s.clients.Load(punchReq.TargetID)
	if err != nil {
		resp.(*PunchResp).Result = RetCode_TargetNotFound
		log.Errorf("cann't find target cliend: %d, err: %+v", punchReq.TargetID, err)
		return
	}

	resp.(*PunchResp).Addr = targetClient.UDPAddr.String()

	getPunchReq := &GetPunchReq{Addr: client.UDPAddr.String()}
	getPunchReqBytes, err := proto.Marshal(getPunchReq)
	if err != nil {
		resp.(*PunchResp).Result = RetCode_ServerError
		log.Errorf("marshal err: %+v", err)
		return
	}

	if n, err := ctx.Conn.WriteToUDP(getPunchReqBytes, targetClient.UDPAddr); err != nil || n != len(getPunchReqBytes) {
		resp.(*PunchResp).Result = RetCode_ServerError
		log.Errorf("write udp err: %+v, client: %+v", err, targetClient)
		return
	}

	return
}
