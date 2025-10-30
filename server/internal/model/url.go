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

func NewUrl(originalUrl, shortCode string) *Url {
	return &Url{
		ShortCode:   shortCode,
		OriginalUrl: originalUrl,
	}
}
