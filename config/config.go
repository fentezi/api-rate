package config

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	Host string `env:"HOST" envDefault:"localhost"`
	Port string `env:"PORT" envDefault:"8080"`

	DbURL string `env:"DB_SOURCE,required"`

	EmailUser string `env:"EMAIL_USER,required"`
	EmailPass string `env:"EMAIL_PASS,required"`
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("unable to load .env file: %v", err)
	}

	cfg := Config{}
	err = env.Parse(&cfg)
	if err != nil {
		log.Fatalf("unable to parse env: %v", err)
	}

	return &cfg

}
