package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Bromolima/url-shortner-go/internal/model"
	"github.com/Bromolima/url-shortner-go/internal/repository"
	"github.com/Bromolima/url-shortner-go/internal/utils"
)

type UrlService interface {
	ShortenUrl(ctx context.Context, originalUrl string) (string, error)
	Redirect(ctx context.Context, shortCode string) (string, error)
}

type urlService struct {
	repository repository.UrlRepository
}

func NewUrlService(repository repository.UrlRepository) UrlService {
	return &urlService{
		repository: repository,
	}
}

func (s *urlService) ShortenUrl(ctx context.Context, originalUrl string) (string, error) {
	shortCode, err := s.repository.FindByOriginalUrl(ctx, originalUrl)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return "", err
	}

	if shortCode != "" {
		return shortCode, nil
	}

	shortCode = utils.GenerateShortCode(utils.ShortCodeLength)

	url := model.NewUrl(originalUrl, shortCode)
	if err := s.repository.Save(ctx, url); err != nil {
		return "", err
	}

	return shortCode, nil
}

func (s *urlService) Redirect(ctx context.Context, shortCode string) (string, error) {
	originalUrl, err := s.repository.FindByShortCode(ctx, shortCode)
	if err != nil {
		return "", err
	}

	return originalUrl, nil
}
