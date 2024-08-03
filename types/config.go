package types

import (
	"log"
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

func (cfg *InitialConfig) ParseFromFile() {
	f, err := os.Open("config.yml")

	if err != nil {
		processError(err)
	}

	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		processError(err)
	}
}

func (cfg *InitialConfig) ParseFromEnv() {
	err := envconfig.Process("", cfg)
	if err != nil {
		processError(err)
	}
}

func processError(err error) {
	log.Fatal(err)
}
