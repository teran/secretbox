package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"

	secretsRepository "github.com/teran/secretbox/repository/secrets/mock"
	tokensRepository "github.com/teran/secretbox/repository/tokens/mock"
)

func (s *serviceTestSuite) TestCreateSecret() {
	s.secretsRepoMock.On("Create", "secret_name", "secret_value").Return(nil).Once()

	err := s.svc.CreateSecret(s.ctx, "secret_name", "secret_value")
	s.Require().NoError(err)
}

func (s *serviceTestSuite) TestCreateToken() {
	s.tokensRepoMock.On("Create", "secret_name", "some_token", defaultTTL).Return(nil).Once()

	err := s.svc.CreateToken(s.ctx, "secret_name", "some_token")
	s.Require().NoError(err)
}

func (s *serviceTestSuite) TestGetSecret() {
	s.tokensRepoMock.On("Redeem", "secret_name", "some_token").Return(nil).Once()
	s.secretsRepoMock.On("Get", "secret_name").Return("some_secret", nil).Once()

	secret, err := s.svc.GetSecret(s.ctx, "secret_name", "some_token")
	s.Require().NoError(err)
	s.Require().Equal("some_secret", secret)
}

func (s *serviceTestSuite) TestGetSecretNoAuth() {
	s.secretsRepoMock.On("Get", "secret_name").Return("some_secret", nil).Once()

	secret, err := s.svc.GetSecretNoAuth(s.ctx, "secret_name")
	s.Require().NoError(err)
	s.Require().Equal("some_secret", secret)
}

func (s *serviceTestSuite) TestIsSecretRegistered() {
	s.secretsRepoMock.On("IsPresent", "existent_secret").Return(true, nil).Once()
	s.secretsRepoMock.On("IsPresent", "not_existent_secret").Return(false, nil).Once()

	ok, err := s.svc.IsSecretRegistered(s.ctx, "existent_secret")
	s.Require().NoError(err)
	s.Require().True(ok)

	ok, err = s.svc.IsSecretRegistered(s.ctx, "not_existent_secret")
	s.Require().NoError(err)
	s.Require().False(ok)
}

// ========================================================================
// Test suite setup
// ========================================================================
type serviceTestSuite struct {
	suite.Suite

	ctx context.Context
	svc Service

	secretsRepoMock *secretsRepository.Mock
	tokensRepoMock  *tokensRepository.Mock
}

func (s *serviceTestSuite) SetupTest() {
	s.ctx = context.TODO()

	s.secretsRepoMock = secretsRepository.New()
	s.tokensRepoMock = tokensRepository.New()

	s.svc = New(s.secretsRepoMock, s.tokensRepoMock)
}

func (s *serviceTestSuite) TearDownTest() {
	s.tokensRepoMock.AssertExpectations(s.T())
	s.secretsRepoMock.AssertExpectations(s.T())
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, &serviceTestSuite{})
}
