package model

import (
	"errors"
)

var (
	ErrUrlNotFound = errors.New("url not found")
)

type Url struct {
	ID          int
	OriginalUrl string
}

func NewUrl(originalUrl string) *Url {
	return &Url{
		OriginalUrl: originalUrl,
	}
}
