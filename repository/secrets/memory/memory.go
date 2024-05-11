package memory

import (
	"context"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/teran/secretbox/repository/secrets"
)

type memory struct {
	data  map[string]string
	mutex *sync.Mutex
}

func New() secrets.Repository {
	return &memory{
		data:  make(map[string]string),
		mutex: &sync.Mutex{},
	}
}

func (m *memory) Create(_ context.Context, name, value string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	log.WithFields(log.Fields{
		"component": "secrets_repository",
		"name":      name,
	}).Debug("create secret")

	m.data[name] = value

	return nil
}

func (m *memory) Get(_ context.Context, name string) (string, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	log.WithFields(log.Fields{
		"component": "secret request",
		"name":      name,
	}).Debug("get secret")

	if v, ok := m.data[name]; ok {
		return v, nil
	}
	return "", secrets.ErrNotFound
}

func (m *memory) IsPresent(_ context.Context, name string) (bool, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	_, ok := m.data[name]
	return ok, nil
}
