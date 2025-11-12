package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Bromolima/url-shortner-go/database"
	"github.com/Bromolima/url-shortner-go/internal/model"
	"github.com/jackc/pgx/v5/pgconn"
)

var (
	ErrDuplicateKey = errors.New("duplicate key")
)

//go:generate mockgen -source=url.go -destination=../../mocks/url_repository.go -package=mocks
type UrlRepository interface {
	Save(ctx context.Context, originalUrl string) (int, error)
	Find(ctx context.Context, id int) (string, error)
	FindByOriginalUrl(ctx context.Context, originalUrl string) (int, error)
}

type urlRepository struct {
	db *database.Database
}

func NewUrlRepository(db *database.Database) UrlRepository {
	return &urlRepository{
		db: db,
	}
}

func (r *urlRepository) Save(ctx context.Context, originalUrl string) (int, error) {
	newUrl := model.NewUrl(originalUrl)
	result := r.db.GormDB.WithContext(ctx).Save(newUrl)

	if result.Error != nil {
		var pgErr *pgconn.PgError
		if errors.As(result.Error, &pgErr) {
			if pgErr.Code == "23505" {
				return 0, ErrDuplicateKey
			}
		}

		return 0, fmt.Errorf("save url: %w", result.Error)
	}

	return newUrl.ID, nil
}

func (r *urlRepository) Find(ctx context.Context, id int) (string, error) {
	var url model.Url
	if err := r.db.GormDB.WithContext(ctx).Where("id = ?", id).First(&url).Error; err != nil {
		return "", fmt.Errorf("find url: %w", err)
	}

	return url.OriginalUrl, nil
}

func (r *urlRepository) FindByOriginalUrl(ctx context.Context, originalUrl string) (int, error) {
	var url model.Url
	if err := r.db.GormDB.WithContext(ctx).Where("original_url = ?", originalUrl).First(&url).Error; err != nil {
		return url.ID, fmt.Errorf("find url by original: %w", err)
	}

	return url.ID, nil
}
