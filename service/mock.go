package service

import (
	"context"

	"github.com/stretchr/testify/mock"
)

var _ Service = (*Mock)(nil)

type Mock struct {
	mock.Mock
}

func NewMock() *Mock {
	return &Mock{}
}

func (m *Mock) CreateSecret(_ context.Context, secretName, tokenValue string) error {
	args := m.Called(secretName, secretName)
	return args.Error(0)
}

func (m *Mock) CreateToken(_ context.Context, secretName, tokenValue string) error {
	args := m.Called(secretName, tokenValue)
	return args.Error(0)
}

func (m *Mock) GetSecret(_ context.Context, secretName, tokenValue string) (string, error) {
	args := m.Called(secretName, tokenValue)
	return args.Get(0).(string), args.Error(1)
}

func (m *Mock) GetSecretNoAuth(_ context.Context, secretName string) (string, error) {
	args := m.Called(secretName)
	return args.Get(0).(string), args.Error(1)
}

func (m *Mock) IsSecretRegistered(_ context.Context, secretName string) (bool, error) {
	args := m.Called(secretName)
	return args.Bool(0), args.Error(1)
}
