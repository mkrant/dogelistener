package sessions

import (
	"errors"
	"fmt"
	"github.com/mkrant/dogelistener/internal/server/api"
	"sort"
)

var count = 1

type Manager struct {
	sessions map[string]*Session
}

func NewManager() *Manager {
	return &Manager{sessions: map[string]*Session{}}
}

func (m *Manager) StartSession(stream api.DogeServer_ConnectServer) (*Session, error) {
	id := fmt.Sprintf("%d", count)
	//count++

	sess := NewSession(id, stream)
	m.sessions[id] = sess
	return sess, nil
}

func (m *Manager) StopSession(id string) error {
	sess, ok := m.sessions[id]
	if !ok {
		return errors.New("session not found")
	}

	if err := sess.StopRun(); err != nil {
		return fmt.Errorf("stopping run: %w", err)
	}

	delete(m.sessions, id)

	return nil
}

func (m *Manager) Sessions() []*Session {
	ss := make([]*Session, 0, len(m.sessions))

	for _, s := range m.sessions {
		ss = append(ss, s)
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].id < ss[j].id
	})

	return ss
}

func (m *Manager) GetSession(id string) (*Session, bool) {
	sess, ok := m.sessions[id]
	return sess, ok
}
