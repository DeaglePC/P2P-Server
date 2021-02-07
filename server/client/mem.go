package client

import (
	"fmt"
	"sync"
	"time"
)

type MemStore struct {
	clients *sync.Map
}

func NewClientWithMemStore() *MemStore {
	return &MemStore{clients: new(sync.Map)}
}

func (m *MemStore) Store(info *Info) error {
	m.clients.Store(info.ID, info)
	return nil
}

func (m *MemStore) Load(id uint64) (*Info, error) {
	client, ok := m.clients.Load(id)
	if !ok {
		return nil, fmt.Errorf("can't found client: %d", id)
	}
	return client.(*Info), nil
}

func (m *MemStore) Delete(id uint64) error {
	_, ok := m.clients.Load(id)
	if !ok {
		return fmt.Errorf("can't found client: %d", id)
	}
	m.clients.Delete(id)
	return nil
}

func (m *MemStore) UpdateHeartbeat(id uint64) error {
	client, ok := m.clients.Load(id)
	if !ok {
		return fmt.Errorf("can't found client: %d", id)
	}

	client.(*Info).LastHeartbeatTime = time.Now().Unix()
	return nil
}
