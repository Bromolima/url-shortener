package repository

import (
	"context"

	"github.com/Bromolima/url-shortner-go/database"
	"github.com/Bromolima/url-shortner-go/internal/model"
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
		return 0, result.Error
	}

	return newUrl.ID, nil
}

func (r *urlRepository) Find(ctx context.Context, id int) (string, error) {
	var url model.Url
	if err := r.db.GormDB.WithContext(ctx).Where("id = ?", id).First(&url).Error; err != nil {
		return "", err
	}

	return url.OriginalUrl, nil
}

func (r *urlRepository) FindByOriginalUrl(ctx context.Context, originalUrl string) (int, error) {
	var url model.Url
	if err := r.db.GormDB.WithContext(ctx).Where("original_url = ?", originalUrl).First(&url).Error; err != nil {
		return url.ID, err
	}

	return url.ID, nil
}
