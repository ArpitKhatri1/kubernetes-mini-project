package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB DBConfig
}

type DBConfig struct {
	User     string
	Password string
	Name     string
	SSLMode  string
	Port     string
	Host     string
}

// panic is used for unrecoverable errors
func LoadConfig() *Config {
	_ = godotenv.Load(".env.development")

	return &Config{
		// creating a struct using its [name]{fields:values}
		DB: DBConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     mustGetEnv("POSTGRES_USER"),
			Password: mustGetEnv("POSTGRES_PASSWORD"),
			Name:     mustGetEnv("POSTGRES_DB"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
	}
}

func getEnv(curr string, fallback string) string {
	val, ok := os.LookupEnv(curr)
	if !ok {
		return fallback
	}
	return val
}

func mustGetEnv(curr string) string {
	val, ok := os.LookupEnv(curr)
	if !ok {
		panic(fmt.Sprintf("can't find the env %s", curr))
	}
	return val
}
