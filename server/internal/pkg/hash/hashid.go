package hash

import (
	"github.com/Bromolima/url-shortner-go/config"
	"github.com/speps/go-hashids/v2"
)

const (
	shortCodeLength = 7
	charset         = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func NewHashID() *hashids.HashID {
	hd := hashids.NewData()
	hd.Salt = config.Env.SecretKey
	hd.MinLength = shortCodeLength
	hd.Alphabet = charset
	h, _ := hashids.NewWithData(hd)
	return h
}
