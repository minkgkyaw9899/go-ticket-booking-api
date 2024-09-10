package config

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type EnvConfig struct {
	PORT        string `env:"SERVER_PORT,required"`
	DB_USER     string `env:"DB_USER,required"`
	DB_PASSWORD string `env:"DB_PASSWORD,required"`
	DB_NAME     string `env:"DB_NAME,required"`
	DB_HOST     string `env:"DB_HOST,required"`
	DB_SSL_MODE string `env:"DB_SSL_MODE,required"`
}

func LoadEnvConfig() (*EnvConfig, error) {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Unable to load .env file: %s", err.Error())
		return nil, err
	}

	cfg := &EnvConfig{}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
