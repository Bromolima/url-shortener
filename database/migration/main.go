package main

import (
	"log"
	"log/slog"

	"github.com/Bromolima/url-shortner-go/config"
	"github.com/Bromolima/url-shortner-go/database"
)

func main() {
	err := config.LoadEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.InitPostgres()

	query := `
	CREATE TABLE IF NOT EXISTS urls (
		id CHAR(36) PRIMARY KEY,
		original_url TEXT NOT NULL,
		short_code VARCHAR(10) NOT NULL UNIQUE
	);
	`

	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Database migrated successfully")
}
