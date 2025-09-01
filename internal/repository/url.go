package repository

import (
	"context"
	"database/sql"

	"github.com/Bromolima/url-shortner-go/internal/model"
)

type UrlRepository interface {
	Save(ctx context.Context, url *model.Url) error
	FindByShortCode(ctx context.Context, shortCode string) (string, error)
	FindByOriginalUrl(ctx context.Context, originalUrl string) (string, error)
}

type urlRepository struct {
	db *sql.DB
}

func NewUrlRepository(db *sql.DB) UrlRepository {
	return &urlRepository{
		db: db,
	}
}

func (r *urlRepository) Save(ctx context.Context, url *model.Url) error {
	query := `
		INSERT INTO urls (id, original_url, short_code) 
		VALUES ($1, $2, $3)
	`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, url.ID, url.OriginalUrl, url.ShortCode)
	if err != nil {
		return err
	}

	return nil
}

func (r *urlRepository) FindByShortCode(ctx context.Context, shortCode string) (string, error) {
	query := `
		SELECT original_url 
		FROM urls
		WHERE short_code = $1
	`

	var originalUrl string
	err := r.db.QueryRowContext(ctx, query, shortCode).Scan(&originalUrl)
	if err != nil {
		return "", err
	}

	return originalUrl, nil
}

func (r *urlRepository) FindByOriginalUrl(ctx context.Context, originalUrl string) (string, error) {
	query := `
		SELECT short_code 
		FROM urls 
		WHERE original_url = $1
	`

	var shortCode string
	err := r.db.QueryRowContext(ctx, query, originalUrl).Scan(&shortCode)
	if err != nil {
		return "", err
	}

	return shortCode, nil
}
