package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	// API REST
	ApiRestPort int32 `envconfig:"API_REST_PORT"`

	// EMAIL
	EmailHost     string `envconfig:"EMAIL_HOST"`
	EmailPort     int32  `envconfig:"EMAIL_PORT"`
	EmailUser     string `envconfig:"EMAIL_USER"`
	EmailPassword string `envconfig:"EMAIL_PASSWORD"`
}

var Env Config = Config{
	// API REST
	ApiRestPort: 8080,

	// EMAIL
	EmailHost:     "localhost",
	EmailPort:     583,
	EmailUser:     "user",
	EmailPassword: "password",
}

func InitConfig() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	return envconfig.Process("", &Env)
}
