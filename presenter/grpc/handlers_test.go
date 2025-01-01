//go:build grpc

package grpc

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	grpctest "github.com/teran/go-grpctest"

	proto "github.com/teran/secretbox/presenter/grpc/proto/v1"
	"github.com/teran/secretbox/service"
)

func (s *handlersTestSuite) TestGetSecret() {
	s.serviceMock.On("GetSecret", "test_name", "test_token").Return("some_secret", nil).Once()
	resp, err := s.client.GetSecret(s.ctx, &proto.GetSecretRequest{
		Name:  "test_name",
		Token: "test_token",
	})
	s.Require().NoError(err)
	s.Require().NotNil(resp)
	s.Require().Equal("some_secret", resp.GetSecret())
}

// TODO: add tests w/ error responses

// ========================================================================
// Test suite setup
// ========================================================================
type handlersTestSuite struct {
	suite.Suite

	serviceMock *service.Mock

	srv    grpctest.Server
	ctx    context.Context
	cancel context.CancelFunc

	client   proto.SecretBoxServiceClient
	handlers Handlers
}

func (s *handlersTestSuite) SetupTest() {
	s.ctx, s.cancel = context.WithTimeout(context.Background(), 30*time.Second)

	s.serviceMock = service.NewMock()

	s.handlers = New(s.serviceMock)

	s.srv = grpctest.New()
	s.handlers.Register(s.srv.Server())

	err := s.srv.Run()
	s.Require().NoError(err)

	dial, err := s.srv.DialContext(s.ctx)
	s.Require().NoError(err)

	s.client = proto.NewSecretBoxServiceClient(dial)
}

func (s *handlersTestSuite) TearDownTest() {
	s.srv.Close()
	s.cancel()
}

func TestHandlersTestSuite(t *testing.T) {
	suite.Run(t, &handlersTestSuite{})
}
