package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

var Env Environment

func LoadEnvironment() error {
	err := godotenv.Load()
	if err != nil {

		return err
	}

	Env = Environment{
		Env:       getEnv("ENV", "development"),
		ApiPort:   getEnv("API_PORT", "8080"),
		ApiUrl:    getEnv("API_URL", "http://localhost:8080/"),
		SecretKey: getEnv("SECRET_KEY", "my_secret_key"),
		DB: Postgres{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "admin"),
			Password: getEnv("DB_PASSWORD", "admin"),
			Name:     getEnv("DB_NAME", "url_shortener"),
		},
	}

	slog.Info("Environment loaded")

	return nil
}

func getEnv(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return defaultValue
}
