package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/Bromolima/url-shortner-go/internal/http/responses"
	"github.com/Bromolima/url-shortner-go/internal/model"
	"github.com/Bromolima/url-shortner-go/internal/service"
)

type UrlHandler interface {
	ShortenUrl(w http.ResponseWriter, r *http.Request)
	Redirect(w http.ResponseWriter, r *http.Request)
}

type urlHandler struct {
	service service.UrlService
}

func NewUrlHandler(service service.UrlService) UrlHandler {
	return &urlHandler{
		service: service,
	}
}

func (h *urlHandler) ShortenUrl(w http.ResponseWriter, r *http.Request) {
	slog.Info("Received request to shorten URL")
	var Url model.UrlPayload
	if err := json.NewDecoder(r.Body).Decode(&Url); err != nil {
		slog.Warn("Failed to decode request body", "error", err)
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}
	slog.Info("Request body decoded successfully")

	if _, err := url.ParseRequestURI(Url.OriginalUrl); err != nil {
		slog.Warn("Invalid URL format", "error", err)
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	shortCode, err := h.service.ShortenUrl(r.Context(), Url.OriginalUrl)
	if err != nil {
		slog.Error("Failed to shorten URL", "error", err)
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	slog.Info("URL shortened")

	response := model.UrlResponse{
		ShortCode: fmt.Sprintf("%s/%s", r.Host, shortCode),
	}

	responses.JSON(w, http.StatusCreated, response)
}

func (h *urlHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	shortCode := r.URL.Path[1:]
	originalUrl, err := h.service.Redirect(r.Context(), shortCode)
	if err != nil {
		if errors.Is(err, model.ErrUrlNotFound) {
			slog.Warn("Short code not found", "error", err)
			responses.Err(w, http.StatusNotFound, err)
			return
		}
		slog.Error("Failed to redirect to original url", "error", err)
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	slog.Info("Redirecting to original URL")
	http.Redirect(w, r, originalUrl, http.StatusFound)
}
