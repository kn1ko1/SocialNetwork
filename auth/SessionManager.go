package auth

import (
	"errors"
	"log"
	"socialnetwork/models"
	"sync"
	"time"
)

type SessionManager struct {
	CookieMap map[string]models.User
	lifetime  int
	mtx       sync.Mutex
}

func NewSessionManager() *SessionManager {
	ret := new(SessionManager)
	ret.CookieMap = make(map[string]models.User)
	ret.lifetime = 1800
	return ret
}

func (m *SessionManager) Add(k string, v models.User) error {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	_, exists := m.CookieMap[k]
	if exists {
		return errors.New("session already exists")
	}
	m.CookieMap[k] = v
	log.Printf("Added Session Cookie. Active Sessions: %d\n", len(m.CookieMap))
	go m.DeleteUponExpiry(k)
	return nil
}

func (m *SessionManager) Delete(k string) {
	m.mtx.Lock()
	delete(m.CookieMap, k)
	log.Printf("Deleted Session Cookie. Active Sessions: %d\n", len(m.CookieMap))
	m.mtx.Unlock()
}

func (m *SessionManager) DeleteUponExpiry(k string) {
	<-time.After(time.Duration(m.lifetime) * time.Second)
	m.Delete(k)
}

func (m *SessionManager) Get(k string) (models.User, error) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	ret, exists := m.CookieMap[k]
	if !exists {
		return ret, errors.New("invalid token")
	}
	return ret, nil
}

func (m *SessionManager) Lifetime() int {
	return m.lifetime
}
