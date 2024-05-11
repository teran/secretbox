package secrets

import (
	"context"

	"github.com/pkg/errors"
)

type Repository interface {
	Create(ctx context.Context, name, value string) error
	Get(ctx context.Context, name string) (string, error)
	IsPresent(ctx context.Context, name string) (bool, error)
}

var ErrNotFound = errors.New("not found")
