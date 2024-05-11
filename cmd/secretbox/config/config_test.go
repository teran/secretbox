package config

import (
	"testing"

	"github.com/stretchr/testify/require"

	onepassword "github.com/teran/go-onepassword-cli"
)

func TestNewConfigFromFile(t *testing.T) {
	r := require.New(t)

	cfg, err := NewFromFile("testdata/config.yaml")
	r.NoError(err)
	r.Equal(&Config{
		Server: ServerConfig{
			Protocol: "unix",
			Socket:   "/tmp/secretbox.sock",
		},
		Secrets: []SecretConfig{
			{
				Name:   "secret1",
				Source: "onepassword",
				Label:  "test:SecretKey",
				Kind:   onepassword.KindPassword,
			},
		},
	}, cfg)
}
