package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ApiRestPort int32 `envconfig:"API_REST_PORT"`

	EmailHost     string `envconfig:"EMAIL_HOST"`
	EmailPort     int32  `envconfig:"EMAIL_PORT"`
	EmailUser     string `envconfig:"EMAIL_USER"`
	EmailPassword string `envconfig:"EMAIL_PASSWORD"`
}

var Env Config

func InitConfig() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	return envconfig.Process("", &Env)
}
