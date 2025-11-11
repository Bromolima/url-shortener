package config

import (
	"log/slog"

	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
)

var Env Environment

func LoadEnvironment() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	_, err = env.UnmarshalFromEnviron(&Env)
	if err != nil {
		slog.Error("Failed to unmarshal env", "error", err.Error())
		return err
	}

	slog.Info("Environment loaded")

	return nil
}
