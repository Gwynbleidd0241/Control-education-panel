package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	Port     string `env:"PORT" envDefault:"8081"`
	LogLevel string `env:"LOG_LEVEL" envDefault:"debug"`

	DBHost     string `env:"DB_HOST" envDefault:"localhost"`
	DBPort     string `env:"DB_PORT" envDefault:"5432"`
	DBUser     string `env:"DB_USER" envDefault:"postgres"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"postgres"`
	DBName     string `env:"DB_NAME" envDefault:"prdb"`
	DBSSLMode  string `env:"DB_SSLMODE" envDefault:"disable"`
}

func MustLoad() *Config {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		panic(fmt.Errorf("config error: %w", err))
	}
	return cfg
}

func (c *Config) DSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.DBUser,
		c.DBPassword,
		c.DBHost,
		c.DBPort,
		c.DBName,
		c.DBSSLMode,
	)
}

func (c *Config) Addr() string {
	return ":" + c.Port
}
