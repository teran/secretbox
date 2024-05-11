package main

import (
	"context"
	"fmt"
	"net"
	"os"

	arg "github.com/alexflint/go-arg"
	log "github.com/sirupsen/logrus"
	grpcServer "google.golang.org/grpc"

	onepassword "github.com/teran/go-onepassword-cli"
	"github.com/teran/secretbox/cmd/secretbox/config"
	"github.com/teran/secretbox/presenter/grpc"
	secretsRepository "github.com/teran/secretbox/repository/secrets/memory"
	tokensRepository "github.com/teran/secretbox/repository/tokens/memory"
	"github.com/teran/secretbox/service"
)

type spec struct {
	Config string `arg:"-c,env:SECRETBOX_CONFIG" default:"/etc/secretbox.yaml"`
}

func init() {
	log.SetLevel(log.TraceLevel)
}

func main() {
	ctx := context.TODO()

	lCfg := spec{}
	arg.MustParse(&lCfg)

	log.Trace("starting secretbox service ...")

	cfg, err := config.NewFromFile(lCfg.Config)
	if err != nil {
		panic(err)
	}

	secretsRepo := secretsRepository.New()
	tokensRepo := tokensRepository.New()

	svc := service.New(secretsRepo, tokensRepo)

	for _, secret := range cfg.Secrets {
		switch secret.Source {
		case "onepassword":
			op := onepassword.New()
			v, err := op.GetByLabel(ctx, secret.Kind, secret.Label)
			if err != nil {
				panic(err)
			}

			err = svc.CreateSecret(ctx, secret.Name, v)
			if err != nil {
				panic(err)
			}
		default:
			panic(fmt.Sprintf("unexpected secret source: `%s`", secret.Source))
		}
	}

	presenter := grpc.New(svc)

	gs := grpcServer.NewServer()
	presenter.Register(gs)

	if cfg.Server.Protocol == "unix" {
		if _, err := os.Stat(cfg.Server.Socket); !os.IsNotExist(err) {
			if err := os.Remove(cfg.Server.Socket); err != nil {
				panic(err)
			}
		}
	}

	listener, err := net.Listen(cfg.Server.Protocol, cfg.Server.Socket)
	if err != nil {
		panic(err)
	}

	log.Infof("listening at `%s`", cfg.Server.Socket)

	if err := gs.Serve(listener); err != nil {
		panic(err)
	}
}
