package handler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"

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
	var Url model.UrlPayload
	if err := json.NewDecoder(r.Body).Decode(&Url); err != nil {
		slog.Warn("Failed to decode request body", "error", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if _, err := url.ParseRequestURI(Url.OriginalUrl); err != nil {
		slog.Warn("Invalid URL format", "error", err)
		http.Error(w, "Invalid URL format", http.StatusBadRequest)
		return
	}

	shortCode, err := h.service.ShortenUrl(r.Context(), Url.OriginalUrl)
	if err != nil {
		slog.Error("Failed to shorten URL", "error", err)
		http.Error(w, "Failed to shorten URL", http.StatusInternalServerError)
		return
	}

	slog.Info("URL shortened")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(model.UrlResponse{
		ShortCode: fmt.Sprintf("http://%s/%s", r.Host, shortCode),
	})
}

func (h *urlHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	shortCode := r.URL.Path[1:]
	slog.Info("Received redirect request", "short_code", shortCode)
	originalUrl, err := h.service.Redirect(r.Context(), shortCode)
	if err != nil {
		slog.Warn("Short code not found", "error", err)
		http.Error(w, "Short code not found", http.StatusNotFound)
		return
	}

	slog.Info("Redirecting to original URL")
	http.Redirect(w, r, originalUrl, http.StatusFound)
}
