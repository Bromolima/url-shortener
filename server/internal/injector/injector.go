package injector

import (
	"database/sql"
	"log/slog"

	"github.com/Bromolima/url-shortner-go/internal/http/handler"
	"github.com/Bromolima/url-shortner-go/internal/repository"
	"github.com/Bromolima/url-shortner-go/internal/service"
	"go.uber.org/dig"
)

func SetupInjections(db *sql.DB, c *dig.Container) {
	c.Provide(func() *sql.DB {
		return db
	})
	c.Provide(repository.NewUrlRepository)
	c.Provide(service.NewUrlService)
	c.Provide(handler.NewUrlHandler)
	slog.Info("Injections setup completed")
}
