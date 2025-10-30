package service

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/Bromolima/url-shortner-go/internal/model"
	"github.com/Bromolima/url-shortner-go/internal/repository"
)

//go:generate mockgen -source=url.go -destination=../../mocks/url_service.go -package=mocks
type UrlService interface {
	ShortenUrl(ctx context.Context, originalUrl string) (string, error)
	Redirect(ctx context.Context, shortCode string) (string, error)
}

type urlService struct {
	repository     repository.UrlRepository
	hashUrlService HashUrlService
}

func NewUrlService(repository repository.UrlRepository, hashUrlService HashUrlService) UrlService {
	return &urlService{
		repository:     repository,
		hashUrlService: hashUrlService,
	}
}

func (s *urlService) ShortenUrl(ctx context.Context, originalUrl string) (string, error) {
	id, err := s.repository.FindByOriginalUrl(ctx, originalUrl)
	if err != nil {
		return "", err
	}

	if id != 0 {
		shortCode, err := s.hashUrlService.EncodeUrl(id)
		if err != nil {
			return "", err
		}
		return shortCode, nil
	}

	url := model.NewUrl(originalUrl)
	id, err = s.repository.Save(ctx, url.OriginalUrl)
	if err != nil {
		return "", err
	}

	shortCode, err := s.hashUrlService.EncodeUrl(id)
	if err != nil {
		return "", err
	}

	return shortCode, nil
}

func (s *urlService) Redirect(ctx context.Context, shortCode string) (string, error) {
	id, err := s.hashUrlService.DecodeUrl(shortCode)
	if err != nil {
		slog.Error("Failed to decode url", "error", err)
		return "", model.ErrUrlNotFound
	}

	originalUrl, err := s.repository.FindByShortCode(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", model.ErrUrlNotFound
		}

		return "", err
	}

	return originalUrl, nil
}
