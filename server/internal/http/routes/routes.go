package routes

import (
	"github.com/Bromolima/url-shortner-go/internal/http/handler"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func SetupRoutes(r *gin.Engine, c *dig.Container) error {
	return c.Invoke(func(h *handler.UrlHandler) {
		g := r.Group("v1")
		g.POST("/shorten", h.ShortenUrl)
		g.GET("/{shortCode}", h.Redirect)
	})
}
