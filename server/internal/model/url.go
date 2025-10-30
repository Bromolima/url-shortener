package model

import (
	"errors"
)

var (
	ErrUrlNotFound = errors.New("url not found")
)

type Url struct {
	ID          int
	ShortCode   string
	OriginalUrl string
}

type UrlPayload struct {
	OriginalUrl string `json:"url"`
}

type UrlResponse struct {
	ShortCode string `json:"short_code"`
}

func NewUrl(originalUrl, shortCode string) *Url {
	return &Url{
		ShortCode:   shortCode,
		OriginalUrl: originalUrl,
	}
}
