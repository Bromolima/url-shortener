package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/Bromolima/url-shortner-go/config"
	"github.com/Bromolima/url-shortner-go/database"
	"github.com/Bromolima/url-shortner-go/internal/http/handler"
	"github.com/Bromolima/url-shortner-go/internal/http/routes"
	"github.com/Bromolima/url-shortner-go/internal/pkg/injector"
	"github.com/Bromolima/url-shortner-go/internal/repository"
	"github.com/Bromolima/url-shortner-go/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	if err := config.LoadEnvironment(); err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.RedirectTrailingSlash = true
	router.Use(cors.Default())
	container := dig.New()

	injector.Provide(container, database.NewPostgresConnection)
	injector.Provide(container, service.NewHashUrlService)
	injector.Provide(container, repository.NewUrlRepository)
	injector.Provide(container, service.NewUrlService)
	injector.Provide(container, handler.NewUrlHandler)

	if err := routes.SetupRoutes(router, container); err != nil {
		log.Fatal(err)
	}

	logger.Info("App running", "port", config.Env.Server.Port)
	router.Run(fmt.Sprintf(":%s", config.Env.Server.Port))
}
