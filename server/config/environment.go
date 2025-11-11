package config

import "time"

type Environment struct {
	Env       string `env:"ENVIRONMENT"`
	SecretKey string `env:"SECRET_KEY"`
	Server    Server
	DB        PostgresDB
}

type Server struct {
	Port string `env:"API_PORT"`
	Host string `env:"API_HOST"`
}

type PostgresDB struct {
	Port              string        `env:"DB_PORT"`
	Host              string        `env:"DB_HOST"`
	Name              string        `env:"DB_NAME"`
	User              string        `env:"DB_USER"`
	Password          string        `env:"DB_PASSWORD"`
	Timeout           time.Duration `env:"DB_TIMEOUT"`
	ConnectionTimeout time.Duration `env:"DB_CONNECTION_TIMEOUT"`
}
