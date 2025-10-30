package repository

import (
	"context"
	"database/sql"
	"errors"
)

type UrlRepository interface {
	Save(ctx context.Context, originalUrl string) (int, error)
	FindByShortCode(ctx context.Context, id int) (string, error)
	FindByOriginalUrl(ctx context.Context, originalUrl string) (int, error)
}

type urlRepository struct {
	db *sql.DB
}

func NewUrlRepository(db *sql.DB) UrlRepository {
	return &urlRepository{
		db: db,
	}
}

func (r *urlRepository) Save(ctx context.Context, originalUrl string) (int, error) {
	query := `
		INSERT INTO urls (original_url) 
		VALUES ($1)
		RETURNING id
	`

	var id int
	err := r.db.QueryRowContext(ctx, query, originalUrl).Scan(&id)

	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *urlRepository) FindByShortCode(ctx context.Context, id int) (string, error) {
	query := `
		SELECT original_url 
		FROM urls
		WHERE id = $1
	`

	var originalUrl string
	err := r.db.QueryRowContext(ctx, query, id).Scan(&originalUrl)
	if err != nil {
		return "", err
	}

	return originalUrl, nil
}

func (r *urlRepository) FindByOriginalUrl(ctx context.Context, originalUrl string) (int, error) {
	query := `
		SELECT id
		FROM urls 
		WHERE original_url = $1
	`

	var id int
	err := r.db.QueryRowContext(ctx, query, originalUrl).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}

		return 0, err
	}

	return id, nil
}
