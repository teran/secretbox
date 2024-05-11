package memory

import (
	"context"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/teran/secretbox/repository/tokens"
)

var _ tokens.Repository = (*memory)(nil)

type (
	secretName = string
	token      = string
	expiresAt  = time.Time
)

type memory struct {
	data  map[secretName]map[token]expiresAt
	mutex *sync.Mutex
}

func New() tokens.Repository {
	return &memory{
		data:  make(map[secretName]map[token]expiresAt),
		mutex: &sync.Mutex{},
	}
}

func (m *memory) Create(ctx context.Context, secretName string, tokenValue string, ttl time.Duration) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	log.WithFields(log.Fields{
		"component": "tokens_repository",
		"name":      secretName,
	}).Debug("create token")

	_, ok := m.data[secretName]
	if !ok {
		m.data[secretName] = make(map[token]expiresAt)
	}

	m.data[secretName][tokenValue] = time.Now().UTC().Add(ttl)

	return nil
}

func (m *memory) Redeem(ctx context.Context, secretName string, tokenValue string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	log.WithFields(log.Fields{
		"component": "tokens_repository",
		"name":      secretName,
	}).Debug("redeem token")

	s, ok := m.data[secretName]
	if !ok {
		return tokens.ErrNotFound
	}

	v, ok := s[tokenValue]
	if !ok {
		return tokens.ErrNotFound
	}

	delete(m.data[secretName], tokenValue)

	if v.After(time.Now().UTC()) {
		return nil
	}

	return tokens.ErrNotFound
}

func (m *memory) Cleanup(ctx context.Context) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	for secretName, tokens := range m.data {
		for token, expiresAt := range tokens {
			if time.Now().UTC().After(expiresAt) {
				delete(m.data[secretName], token)
			}
		}
	}

	return nil
}
