package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port        string
	Addr        string
	StudentsURL string
	CoursesURL  string
	CertsURL    string
}

func MustLoad() Config {
	port := getEnv("PORT", "8080")

	cfg := Config{
		Port:        port,
		Addr:        fmt.Sprintf(":%s", port),
		StudentsURL: mustEnv("STUDENTS_URL"),
		CoursesURL:  mustEnv("COURSES_URL"),
		CertsURL:    mustEnv("CERTS_URL"),
	}

	return cfg
}

func mustEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic(key + " is required")
	}
	return v
}

func getEnv(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}
