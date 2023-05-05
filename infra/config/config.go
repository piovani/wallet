package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ApiRestPort int32 `envconfig:"API_REST_PORT"`
}

var Env Config

func InitConfig() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	return envconfig.Process("", &Env)
}
