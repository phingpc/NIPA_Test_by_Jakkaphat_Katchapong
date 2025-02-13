package config

import (
	"os"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	APIPORT    string
}

func LoadConfig() Config {
	return Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "nipa_test"),
		DBPassword: getEnv("DB_PASSWORD", "nipa_test123"),
		DBName:     getEnv("DB_NAME", "nipa_test"),
		APIPORT:    getEnv("API_PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	return v
}