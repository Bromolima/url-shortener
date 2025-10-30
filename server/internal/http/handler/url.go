package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/Bromolima/url-shortner-go/internal/http/dto"
	resterrors "github.com/Bromolima/url-shortner-go/internal/http/rest_errors"
	"github.com/Bromolima/url-shortner-go/internal/model"
	"github.com/Bromolima/url-shortner-go/internal/service"
	"github.com/gin-gonic/gin"
)

type UrlHandler struct {
	service service.UrlService
}

func NewUrlHandler(service service.UrlService) *UrlHandler {
	return &UrlHandler{
		service: service,
	}
}

func (h *UrlHandler) ShortenUrl(c *gin.Context) {
	slog.Info("Received request to shorten URL")
	var payload dto.ShortenUrlPayload
	if err := json.NewDecoder(c.Request.Body).Decode(&payload); err != nil {
		slog.Warn("Failed to decode request body", "error", err)
		restErr := resterrors.NewUnprocessableEntityError("Failed to process request body")
		c.JSON(http.StatusUnprocessableEntity, restErr)
		return
	}

	if _, err := url.ParseRequestURI(payload.OriginalUrl); err != nil {
		slog.Warn("Invalid URL format", "error", err)
		restErr := resterrors.NewBadRequestError("The url is in a invalid format")
		c.JSON(http.StatusBadRequest, restErr)
		return
	}

	shortCode, err := h.service.ShortenUrl(c.Request.Context(), payload.OriginalUrl)
	if err != nil {
		slog.Error("Failed to shorten URL", "error", err)
		restErr := resterrors.NewInternalServerError("An unexpected internal server error ocurred")
		c.JSON(http.StatusInternalServerError, restErr)
		return
	}

	slog.Info("URL shortened")

	response := dto.UrlResponse{
		ShortCode: fmt.Sprintf("%s/v1/%s", c.Request.Host, shortCode),
	}

	c.JSON(http.StatusOK, response)
}

func (h *UrlHandler) Redirect(c *gin.Context) {
	var payload dto.RedirectUrlResponse
	if err := c.ShouldBindUri(&payload); err != nil {
		slog.Warn("Failed to decode request uri", "error", err)
		restErr := resterrors.NewUnprocessableEntityError("Failed to process request uri")
		c.JSON(http.StatusUnprocessableEntity, restErr)
		return
	}

	originalUrl, err := h.service.Redirect(c.Request.Context(), payload.ShortCode)
	if err != nil {
		if errors.Is(err, model.ErrUrlNotFound) {
			slog.Warn("Short code not found", "error", err)
			restErr := resterrors.NewNotFoundError("Short code not found")
			c.JSON(http.StatusNotFound, restErr)
			return
		}
		slog.Error("Failed to redirect to original url", "error", err)
		restErr := resterrors.NewInternalServerError("An unexpected internal server error ocurred")
		c.JSON(http.StatusInternalServerError, restErr)
		return
	}

	slog.Info("Redirecting to original URL")
	c.Redirect(http.StatusFound, originalUrl)
}
