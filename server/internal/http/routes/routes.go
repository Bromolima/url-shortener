package routes

import (
	"log/slog"

	"github.com/Bromolima/url-shortner-go/internal/http/handler"
	"go.uber.org/dig"
)

func SetupRoutes(r *Router, c *dig.Container) error {
	return c.Invoke(func(urlHandler handler.UrlHandler) {
		r.POST("v1/shorten", urlHandler.ShortenUrl)
		r.GET("v1/{shortCode}", urlHandler.Redirect)
		slog.Info("Routes setup completed")
	})
}
