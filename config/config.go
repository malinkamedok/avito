package config

import "github.com/caarlos0/env/v6"

type Config struct {
	AppPort     string `env:"APP_PORT" envDefault:"9000"`
	PostgresUrl string `env:"POSTGRES_URL" envDefault:"postgresql://postgres:example@localhost:5432/postgres"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
