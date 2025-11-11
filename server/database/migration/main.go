package main

import (
	"log"

	"github.com/Bromolima/url-shortner-go/config"
	"github.com/Bromolima/url-shortner-go/database"
	"github.com/Bromolima/url-shortner-go/internal/model"
)

func main() {
	err := config.LoadEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.NewPostgresConnection()
	if err != nil {
		log.Fatal(err)
	}

	db.GormDB.AutoMigrate(model.Url{})
}
