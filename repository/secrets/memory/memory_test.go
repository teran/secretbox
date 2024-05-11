package memory

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMemoryRepository(t *testing.T) {
	ctx := context.TODO()
	r := require.New(t)

	const (
		key   = "some_key"
		value = "some value"
	)

	repo := New()

	err := repo.Create(ctx, key, value)
	r.NoError(err)

	ok, err := repo.IsPresent(ctx, key)
	r.NoError(err)
	r.True(ok)

	ok, err = repo.IsPresent(ctx, "unexistent")
	r.NoError(err)
	r.False(ok)

	v, err := repo.Get(ctx, key)
	r.NoError(err)
	r.Equal(value, v)
}
