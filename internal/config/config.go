package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	App      AppConfig
	Log      LogConfig
	Postgres PostgresConfig
}

type AppConfig struct {
	Environment string `env:"APP_ENV" env-required:"true"`
}

type LogConfig struct {
	Level  string `env:"LOG_LEVEL" env-required:"true"`
	Format string `env:"LOG_FORMAT" env-required:"true"`
}

type PostgresConfig struct {
	Host     string `env:"DB_HOST" env-required:"true"`
	Port     uint16 `env:"DB_PORT" env-required:"true"`
	Name     string `env:"DB_NAME" env-required:"true"`
	User     string `env:"DB_USER" env-required:"true"`
	Password string `env:"DB_PASSWORD" env-required:"true"`
	SSLMode  string `env:"DB_SSLMODE" env-required:"true"`
}

func Load() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("load .env: %w", err)
	}

	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, fmt.Errorf("read environment: %w", err)
	}

	return &cfg, nil
}
