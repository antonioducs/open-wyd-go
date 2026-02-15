package session

import (
	"net"
	"sync"
	"sync/atomic"
)

type Session struct {
	ID   uint32
	Conn net.Conn
}

type Manager struct {
	sessions map[uint32]*Session
	mu       sync.RWMutex
	counter  uint32
}

var Global = &Manager{
	sessions: make(map[uint32]*Session),
}

func (m *Manager) NextID() uint32 {
	return atomic.AddUint32(&m.counter, 1)
}

func (m *Manager) Add(conn net.Conn) *Session {
	id := m.NextID()
	s := &Session{ID: id, Conn: conn}

	m.mu.Lock()
	m.sessions[id] = s
	m.mu.Unlock()

	return s
}

func (m *Manager) Remove(id uint32) {
	m.mu.Lock()
	delete(m.sessions, id)
	m.mu.Unlock()
}

func (m *Manager) Get(id uint32) *Session {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.sessions[id]
}
