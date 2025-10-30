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
	if err != nil {
		log.Fatal(err)
	}

	query := `
	CREATE TABLE IF NOT EXISTS urls (
		id SERIAL PRIMARY KEY,
		original_url TEXT NOT NULL
	);
	`

	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Database migrated successfully")
}
