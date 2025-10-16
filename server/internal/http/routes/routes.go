package routes

import (
	"log/slog"

	"github.com/Bromolima/url-shortner-go/internal/http/handler"
	"go.uber.org/dig"
)

func SetupRoutes(r *Router, c *dig.Container) error {
	return c.Invoke(func(urlHandler handler.UrlHandler) {
		r.POST("/shorten", urlHandler.ShortenUrl)
		r.GET("/{shortCode}", urlHandler.Redirect)
		slog.Info("Routes setup completed")
	})
}
