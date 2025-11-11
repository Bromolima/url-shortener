package injector

import (
	"log/slog"

	"github.com/Bromolima/url-shortner-go/database"
	"github.com/Bromolima/url-shortner-go/internal/http/handler"
	"github.com/Bromolima/url-shortner-go/internal/repository"
	"github.com/Bromolima/url-shortner-go/internal/service"
	"go.uber.org/dig"
)

func SetupInjections(db *database.Database, c *dig.Container) {
	c.Provide(func() *database.Database {
		return db
	})
	c.Provide(service.NewHashUrlService)
	c.Provide(repository.NewUrlRepository)
	c.Provide(service.NewUrlService)
	c.Provide(handler.NewUrlHandler)
	slog.Info("Injections setup completed")
}
