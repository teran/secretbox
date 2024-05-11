package mock

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}

func New() *Mock {
	return &Mock{}
}

func (m *Mock) Create(_ context.Context, name, value string) error {
	args := m.Called(name, value)
	return args.Error(0)
}

func (m *Mock) Get(_ context.Context, name string) (string, error) {
	args := m.Called(name)
	return args.Get(0).(string), args.Error(1)
}

func (m *Mock) IsPresent(_ context.Context, name string) (bool, error) {
	args := m.Called(name)
	return args.Bool(0), args.Error(1)
}
