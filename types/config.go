package types

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type InitialConfig struct {
	Server struct {
		Port string `yaml:"port" envconfig:"SERVER_PORT"`
		Host string `yaml:"host" envconfig:"SERVER_HOST"`
	} `yaml:"server"`
	Database struct {
		Dsn string `yaml:"dsn" envconfig:"DSN"`
	} `yaml:"database"`
}

func (cfg *InitialConfig) ParseFromFile() error {
	f, err := os.Open("config.yml")

	if err != nil {
		return err
	}

	defer f.Close()

	decoder := yaml.NewDecoder(f)
	return decoder.Decode(&cfg)
}

func (cfg *InitialConfig) ParseFromEnv() error {
	return envconfig.Process("", cfg)
}
