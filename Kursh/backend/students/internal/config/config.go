package config

import "os"

type Config struct {
	Port        string
	DatabaseURL string
}

func MustLoad() Config {
	port := getenv("PORT", "8080")
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		panic("DATABASE_URL is required")
	}
	return Config{
		Port:        port,
		DatabaseURL: dsn,
	}
}

func getenv(k, def string) string {
	v := os.Getenv(k)
	if v == "" {
		return def
	}
	return v
}
