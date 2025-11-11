package model

import (
	"errors"
)

var (
	ErrUrlNotFound = errors.New("url not found")
)

type Url struct {
	ID          int    `gorm:"primaryKey;autoIncrement"`
	OriginalUrl string `gorm:"type:text"`
}

func NewUrl(originalUrl string) *Url {
	return &Url{
		OriginalUrl: originalUrl,
	}
}
