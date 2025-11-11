package hash

import (
	"fmt"

	"github.com/Bromolima/url-shortner-go/config"
	"github.com/speps/go-hashids/v2"
)

const (
	shortCodeLength = 7
	charset         = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

//go:generate mockgen -source=hashid.go -destination=../../mocks/hashid.go -package=mocks
type IDHasher interface {
	EncodeUrl(id int) (string, error)
	DecodeUrl(shortUrl string) (int, error)
}

type idHasher struct {
	hashID *hashids.HashID
}

func NewIDHasher() (IDHasher, error) {
	hd := hashids.NewData()
	hd.Salt = config.Env.Salt
	hd.MinLength = shortCodeLength
	hd.Alphabet = charset

	h, err := hashids.NewWithData(hd)
	if err != nil {
		return nil, fmt.Errorf("new hash id: %w", err)
	}

	return &idHasher{
		hashID: h,
	}, err
}

func (h *idHasher) EncodeUrl(id int) (string, error) {
	shortCode, err := h.hashID.Encode([]int{id})
	if err != nil {
		return "", err
	}

	return shortCode, nil
}

func (h *idHasher) DecodeUrl(shortUrl string) (int, error) {
	id, err := h.hashID.DecodeWithError(shortUrl)
	if err != nil {
		return 0, err
	}

	if len(id) == 0 {
		return 0, fmt.Errorf("invalid code")
	}

	return id[0], nil
}
