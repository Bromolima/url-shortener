package service

import (
	"errors"

	"github.com/Bromolima/url-shortner-go/internal/pkg/hash"
	"github.com/speps/go-hashids/v2"
)

//go:generate mockgen -source=hash_url.go -destination=../../mocks/hash_url_service.go -package=mocks
type HashUrlService interface {
	EncodeUrl(id int) (string, error)
	DecodeUrl(shortUrl string) (int, error)
}

type hashUrlService struct {
	hasher *hashids.HashID
}

func NewHashUrlService() HashUrlService {
	return &hashUrlService{
		hasher: hash.NewHashID(),
	}
}

func (s *hashUrlService) EncodeUrl(id int) (string, error) {
	shortCode, err := s.hasher.Encode([]int{id})
	if err != nil {
		return "", err
	}

	return shortCode, nil
}

func (s *hashUrlService) DecodeUrl(shortUrl string) (int, error) {
	id, err := s.hasher.DecodeWithError(shortUrl)
	if err != nil {
		return 0, err
	}
	if len(id) == 0 {
		return 0, errors.New("invalid short code")
	}

	return id[0], nil
}
