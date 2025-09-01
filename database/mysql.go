package database

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/Bromolima/url-shortner-go/config"
	_ "github.com/lib/pq"
)

func InitPostgres() (*sql.DB, error) {
	dns := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.Env.DB.User,
		config.Env.DB.Password,
		config.Env.DB.Host,
		config.Env.DB.Port,
		config.Env.DB.Name,
	)

	db, err := sql.Open("postgres", dns)
	if err != nil {
		slog.Error("Error connecting to the database", "error", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		slog.Error("Error pinging the database", "error", err)
		return nil, err
	}

	slog.Info("Database connected successfully")
	return db, nil
}

func Migrate(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS urls (
		id CHAR(36) PRIMARY KEY,
		original_url TEXT NOT NULL,
		short_code VARCHAR(10) NOT NULL UNIQUE
	);
	`

	_, err := db.Exec(query)
	if err != nil {
		slog.Error("Error creating urls table", "error", err)
		return err
	}

	slog.Info("Database migrated successfully")
	return nil
}
