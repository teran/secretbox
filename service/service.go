package service

import (
	"context"
	"time"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/teran/secretbox/repository/secrets"
	"github.com/teran/secretbox/repository/tokens"
)

const defaultTTL = 10 * time.Second

type Service interface {
	CreateSecret(ctx context.Context, name, value string) error
	CreateToken(ctx context.Context, secretName, tokenValue string) error
	GetSecret(ctx context.Context, secretName, tokenValue string) (string, error)
	GetSecretNoAuth(ctx context.Context, secretName string) (string, error)
	IsSecretRegistered(ctx context.Context, secretName string) (bool, error)
}

type service struct {
	secretsRepo secrets.Repository
	tokensRepo  tokens.Repository
}

func New(secretsRepo secrets.Repository, tokensRepo tokens.Repository) Service {
	return &service{
		secretsRepo: secretsRepo,
		tokensRepo:  tokensRepo,
	}
}

func (s *service) CreateSecret(ctx context.Context, name, value string) error {
	log.WithFields(log.Fields{
		"component": "service",
		"name":      name,
	}).Info("creating secret")

	return s.secretsRepo.Create(ctx, name, value)
}

func (s *service) CreateToken(ctx context.Context, secretName, tokenValue string) error {
	log.WithFields(log.Fields{
		"component":   "service",
		"secret_name": secretName,
	}).Info("creating token")

	return s.tokensRepo.Create(ctx, secretName, tokenValue, defaultTTL)
}

func (s *service) GetSecret(ctx context.Context, secretName, tokenValue string) (string, error) {
	log.WithFields(log.Fields{
		"component":   "service",
		"secret_name": secretName,
	}).Info("secret request")

	err := s.tokensRepo.Redeem(ctx, secretName, tokenValue)
	if err != nil {
		return "", errors.Wrap(err, "error verifying authentication token")
	}

	return s.secretsRepo.Get(ctx, secretName)
}

func (s *service) GetSecretNoAuth(ctx context.Context, secretName string) (string, error) {
	log.WithFields(log.Fields{
		"component":   "service",
		"secret_name": secretName,
	}).Info("secret request")

	return s.secretsRepo.Get(ctx, secretName)
}

func (s *service) IsSecretRegistered(ctx context.Context, secretName string) (bool, error) {
	log.WithFields(log.Fields{
		"component":   "service",
		"secret_name": secretName,
	}).Info("secret presence request")

	return s.secretsRepo.IsPresent(ctx, secretName)
}
