package model

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrUrlNotFound = errors.New("url not found")
)

type Url struct {
	ID          string
	OriginalUrl string
	ShortCode   string
}

type UrlPayload struct {
	OriginalUrl string `json:"url"`
}

type UrlResponse struct {
	ShortCode string `json:"short_code"`
}

func NewUrl(originalUrl, shortCode string) *Url {
	return &Url{
		ID:          uuid.New().String(),
		OriginalUrl: originalUrl,
		ShortCode:   shortCode,
	}
}
