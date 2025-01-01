package config

import (
	"encoding/json"
	"os"

	yaml "gopkg.in/yaml.v3"

	onepassword "github.com/teran/go-onepassword-cli"
)

type ServerConfig struct {
	Protocol string `json:"protocol"`
	Socket   string `json:"socket"`
}

type SecretConfig struct {
	Name   string           `json:"name"`
	Source string           `json:"source"`
	Label  string           `json:"label"`
	Kind   onepassword.Kind `json:"kind"`
}

type Config struct {
	Server  ServerConfig   `json:"server"`
	Secrets []SecretConfig `json:"secrets"`
}

func NewFromFile(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	configFileContents := map[string]any{}
	err = yaml.Unmarshal(data, configFileContents)
	if err != nil {
		return nil, err
	}

	configJSON, err := json.Marshal(configFileContents)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = json.Unmarshal(configJSON, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, cfg.Validate()
}

func (c *Config) Validate() error {
	return nil
}
