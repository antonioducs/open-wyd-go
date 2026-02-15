package session

import (
	"errors"
	"sync"

	"github.com/antonioducs/wyd/pkg/domain/entity"
)

type UserSession struct {
	Mu           sync.RWMutex
	Account      *entity.Account
	Characters   entity.CharacterList
	SelectedChar *entity.Character
}

type Manager struct {
	sessions map[uint32]*UserSession
	mu       sync.RWMutex
}

func NewManager() *Manager {
	return &Manager{
		sessions: make(map[uint32]*UserSession),
	}
}

func (m *Manager) Add(sessionID uint32, acc *entity.Account, chars []*entity.Character) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.sessions[sessionID] = &UserSession{
		Account:    acc,
		Characters: chars,
	}
}

func (m *Manager) Get(sessionID uint32) (*UserSession, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	s, ok := m.sessions[sessionID]
	return s, ok
}

func (m *Manager) Remove(sessionID uint32) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.sessions, sessionID)
}

func (s *UserSession) AppendCharacter(newChar *entity.Character) {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	s.Characters = append(s.Characters, newChar)
}

func (s *UserSession) SelectCharacter(index int) error {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	if index < 0 || index >= len(s.Characters) {
		return errors.New("index out of range")
	}
	s.SelectedChar = s.Characters[index]
	return nil
}
