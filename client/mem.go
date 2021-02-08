package client

import (
	"fmt"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

type MemSet struct {
	clients *sync.Map
}

func NewClientWithMemStore() *MemSet {
	return &MemSet{clients: new(sync.Map)}
}

func (m *MemSet) Store(info *Info) error {
	m.clients.Store(info.ID, info)
	return nil
}

func (m *MemSet) Load(id uint64) (*Info, error) {
	client, ok := m.clients.Load(id)
	if !ok {
		return nil, fmt.Errorf("can't found client: %d", id)
	}
	return client.(*Info), nil
}

func (m *MemSet) Delete(id uint64) error {
	_, ok := m.clients.Load(id)
	if !ok {
		return fmt.Errorf("can't found client: %d", id)
	}
	m.clients.Delete(id)
	return nil
}

func (m *MemSet) UpdateHeartbeat(id uint64) error {
	client, ok := m.clients.Load(id)
	if !ok {
		return fmt.Errorf("can't found client: %d", id)
	}

	client.(*Info).LastHeartbeatTime = time.Now().Unix()
	return nil
}

func (m *MemSet) DeleteTimeOut(timeout int64) (err error) {
	m.clients.Range(func(key, value interface{}) bool {
		if time.Now().Unix()-value.(*Info).LastHeartbeatTime > timeout {
			if err = m.Delete(key.(uint64)); err != nil {
				log.Errorf("delete timeout client error: %+v", err)
			}
		}
		return true
	})
	return
}
