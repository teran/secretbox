package memory

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/teran/secretbox/repository/tokens"
)

func TestCreateRedeem(t *testing.T) {
	ctx := context.TODO()
	r := require.New(t)

	const (
		secretName = "some_secret_name"
		tokenValue = "some_token"
	)

	repo := New().(*memory)

	err := repo.Create(ctx, secretName, tokenValue, 1*time.Second)
	r.NoError(err)

	_, ok := repo.data[secretName]
	r.True(ok)

	_, ok = repo.data[secretName][tokenValue]
	r.True(ok)

	err = repo.Redeem(ctx, secretName, tokenValue)
	r.NoError(err)

	err = repo.Redeem(ctx, secretName, tokenValue)
	r.Error(err)
	r.Equal(tokens.ErrNotFound, err)
}

func TestCleanup(t *testing.T) {
	ctx := context.TODO()
	r := require.New(t)

	const (
		secretName = "some_secret_name"
		tokenValue = "some_token"
	)

	repo := New().(*memory)

	err := repo.Create(ctx, secretName, tokenValue, -1*time.Second)
	r.NoError(err)

	_, ok := repo.data[secretName]
	r.True(ok)

	_, ok = repo.data[secretName][tokenValue]
	r.True(ok)

	err = repo.Cleanup(ctx)
	r.NoError(err)

	_, ok = repo.data[secretName]
	r.True(ok)

	_, ok = repo.data[secretName][tokenValue]
	r.False(ok)
}
