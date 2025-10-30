package repository

import (
	"context"
	"database/sql"
	"errors"
)

type UrlRepository interface {
	Save(ctx context.Context, originalUrl string) (int, error)
	FindByShortCode(ctx context.Context, shortCode string) (string, error)
	FindByOriginalUrl(ctx context.Context, originalUrl string) (*string, error)
	SaveShortCode(ctx context.Context, shortCode string, id int) error
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

func (r *urlRepository) FindByOriginalUrl(ctx context.Context, originalUrl string) (*string, error) {
	query := `
		SELECT short_code
		FROM urls 
		WHERE original_url = $1
	`

	var shortCode *string
	err := r.db.QueryRowContext(ctx, query, originalUrl).Scan(&shortCode)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return shortCode, nil
}

func (r *urlRepository) SaveShortCode(ctx context.Context, shortCode string, id int) error {
	query := `
		UPDATE urls SET short_code = $1
		WHERE id = $2
	`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, shortCode, id)
	if err != nil {
		return err
	}

	return nil
}
