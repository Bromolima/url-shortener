package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/Bromolima/url-shortner-go/config"
	"github.com/Bromolima/url-shortner-go/database"
	"github.com/Bromolima/url-shortner-go/internal/http/routes"
	"github.com/Bromolima/url-shortner-go/internal/pkg/injector"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	if err := config.LoadEnvironment(); err != nil {
		log.Fatal(err)
	}

	db, err := database.NewPostgresConnection()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.RedirectTrailingSlash = true
	router.Use(cors.Default())
	c := dig.New()

	injector.SetupInjections(db, c)

	if err := routes.SetupRoutes(router, c); err != nil {
		log.Fatal(err)
	}

	logger.Info("App running", "port", config.Env.ApiPort)
	router.Run(fmt.Sprintf(":%s", config.Env.ApiPort))
}
