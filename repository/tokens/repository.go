package tokens

import (
	"context"
	"time"

	"github.com/pkg/errors"
)

type Repository interface {
	Create(ctx context.Context, secretName string, token string, ttl time.Duration) error
	Redeem(ctx context.Context, secretName string, token string) error
	Cleanup(ctx context.Context) error
}

var ErrNotFound = errors.New("not found")
