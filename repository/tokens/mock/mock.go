package mock

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"

	"github.com/teran/secretbox/repository/tokens"
)

var _ tokens.Repository = (*Mock)(nil)

type Mock struct {
	mock.Mock
}

func New() *Mock {
	return &Mock{}
}

func (m *Mock) Create(_ context.Context, secretName string, token string, ttl time.Duration) error {
	args := m.Called(secretName, token, ttl)
	return args.Error(0)
}

func (m *Mock) Redeem(_ context.Context, secretName string, token string) error {
	args := m.Called(secretName, token)
	return args.Error(0)
}

func (m *Mock) Cleanup(ctx context.Context) error {
	args := m.Called()
	return args.Error(0)
}
