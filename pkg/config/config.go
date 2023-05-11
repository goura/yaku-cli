package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Debug        bool
	DummyAPIKey  string `split_words:"true"`
	DeeplAuthKey string `split_words:"true"` //We don't want DEEP_L so we use Deepl
}

func NewEnvConfig() Config {
	var conf Config
	err := envconfig.Process("yaku", &conf)
	if err != nil {
		log.Fatal(err.Error())
	}
	return conf
}
