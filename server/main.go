package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/Bromolima/url-shortner-go/config"
	"github.com/Bromolima/url-shortner-go/database"
	"github.com/Bromolima/url-shortner-go/internal/http/routes"
	"github.com/Bromolima/url-shortner-go/internal/injector"
	"go.uber.org/dig"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	if err := config.LoadEnvironment(); err != nil {
		log.Fatal(err)
	}

	db, err := database.InitPostgres()
	if err != nil {
		log.Fatal(err)
	}

	if err := database.Migrate(db); err != nil {
		log.Fatal(err)
	}

	router := routes.NewRouter()
	c := dig.New()

	injector.SetupInjections(db, c)

	if err := routes.SetupRoutes(router, c); err != nil {
		log.Fatal(err)
	}

	logger.Info("App running", "port", config.Env.ApiPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.Env.ApiPort), router.Mux))
}
