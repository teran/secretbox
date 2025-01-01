package main

import (
	"context"
	"fmt"
	"net"

	arg "github.com/alexflint/go-arg"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	proto "github.com/teran/secretbox/presenter/grpc/proto/v1"
)

var (
	appVersion      = "n/a (dev build)"
	buildTimestamp  = "n/a (dev build)"
	commitTimestamp = "n/a (dev build)"
	gitCommit       = "n/a (dev build)"
)

type spec struct {
	Protocol     string `arg:"-p,env:LISTEN_PROTOCOL" help:"listen protocol: unix or tcp" default:"unix"`
	ListenSocket string `arg:"-l,env:LISTEN_SOCKET" help:"listen socket: unix domain socket path or TCP port number" default:"/tmp/secretbox.sock"`
	SecretName   string `arg:"-s" help:"secret name to retrieve"`
	AccessToken  string `arg:"-t" help:"access token"`
}

func main() {
	cfg := spec{}
	arg.MustParse(&cfg)

	log.WithFields(log.Fields{
		"version":          appVersion,
		"build_timestamp":  buildTimestamp,
		"commit_timestamp": commitTimestamp,
		"git_commit":       gitCommit,
	}).Trace("running secretbox-cli")

	ctx := context.TODO()

	dialer := func(ctx context.Context, addr string) (net.Conn, error) {
		var d net.Dialer
		return d.DialContext(ctx, cfg.Protocol, cfg.ListenSocket)
	}

	conn, err := grpc.NewClient(
		cfg.ListenSocket,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithContextDialer(dialer))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := proto.NewSecretBoxServiceClient(conn)
	resp, err := client.GetSecret(ctx, &proto.GetSecretRequest{
		Name:  cfg.SecretName,
		Token: cfg.AccessToken,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.GetSecret())
}
