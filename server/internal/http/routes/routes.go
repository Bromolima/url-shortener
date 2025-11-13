package routes

import (
	"github.com/Bromolima/url-shortner-go/internal/http/handler"
	"github.com/Bromolima/url-shortner-go/internal/pkg/injector"
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

func SetupRoutes(r *gin.Engine, c *dig.Container) error {
	urlHandler, err := injector.Resolve[*handler.UrlHandler](c)
	if err != nil {
		return err
	}

	r.POST("/shorten", urlHandler.ShortenUrl)
	r.GET("/:short_code", urlHandler.Redirect)

	return nil
}
