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

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

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
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.Env.ApiPort), enableCORS(router.Mux)))
}
