package service

import (
	"context"
	"errors"
	"testing"

	"github.com/Bromolima/url-shortner-go/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestUrlService_ShortenUrl(t *testing.T) {
	t.Run("should shorten url and return it", func(t *testing.T) {
		control := gomock.NewController(t)
		defer control.Finish()

		ctx := context.Background()
		idHasherMock := mocks.NewMockIDHasher(control)
		urlRepoMock := mocks.NewMockUrlRepository(control)
		service := NewUrlService(urlRepoMock, idHasherMock)

		originalUrl := "http://example.com"
		existingID := 123
		ExpectedshortCode := "abcdefg"

		urlRepoMock.EXPECT().FindByOriginalUrl(ctx, originalUrl).Return(0, gorm.ErrRecordNotFound)
		urlRepoMock.EXPECT().Save(ctx, originalUrl).Return(existingID, nil)
		idHasherMock.EXPECT().EncodeUrl(existingID).Return(ExpectedshortCode, nil)

		shortCode, err := service.ShortenUrl(ctx, originalUrl)

		assert.NoError(t, err)
		assert.Equal(t, shortCode, ExpectedshortCode)
	})

	t.Run("should find an already saved url and return it shortened", func(t *testing.T) {
		control := gomock.NewController(t)
		defer control.Finish()

		ctx := context.Background()
		idHasherMock := mocks.NewMockIDHasher(control)
		urlRepoMock := mocks.NewMockUrlRepository(control)
		service := NewUrlService(urlRepoMock, idHasherMock)

		originalUrl := "http://example.com"
		existingID := 123
		ExpectedshortCode := "abcdefg"

		urlRepoMock.EXPECT().FindByOriginalUrl(ctx, originalUrl).Return(123, nil)
		idHasherMock.EXPECT().EncodeUrl(existingID).Return(ExpectedshortCode, nil)

		shortCode, err := service.ShortenUrl(ctx, originalUrl)

		assert.NoError(t, err)
		assert.Equal(t, shortCode, ExpectedshortCode)
	})

	t.Run("should return an error when repository fails to find url", func(t *testing.T) {
		control := gomock.NewController(t)
		defer control.Finish()

		ctx := context.Background()
		idHasherMock := mocks.NewMockIDHasher(control)
		urlRepoMock := mocks.NewMockUrlRepository(control)
		service := NewUrlService(urlRepoMock, idHasherMock)

		originalUrl := "http://example.com"
		expecterError := errors.New("database connection failed")

		urlRepoMock.EXPECT().FindByOriginalUrl(ctx, originalUrl).Return(0, expecterError)

		shortCode, err := service.ShortenUrl(ctx, originalUrl)

		assert.Error(t, err)
		assert.Equal(t, expecterError, err)
		assert.Equal(t, "", shortCode)
	})

	t.Run("should return an error when repository fails to save url", func(t *testing.T) {
		control := gomock.NewController(t)
		defer control.Finish()

		ctx := context.Background()
		idHasherMock := mocks.NewMockIDHasher(control)
		urlRepoMock := mocks.NewMockUrlRepository(control)
		service := NewUrlService(urlRepoMock, idHasherMock)

		originalUrl := "http://example.com"
		expecterError := errors.New("database connection failed")

		urlRepoMock.EXPECT().FindByOriginalUrl(ctx, originalUrl).Return(0, nil)
		urlRepoMock.EXPECT().Save(ctx, originalUrl).Return(0, expecterError)

		shortCode, err := service.ShortenUrl(ctx, originalUrl)

		assert.Error(t, err)
		assert.Equal(t, expecterError, err)
		assert.Equal(t, "", shortCode)
	})

	t.Run("should return an error when repository finds an existing url, but hasher fails to encode id", func(t *testing.T) {
		control := gomock.NewController(t)
		defer control.Finish()

		ctx := context.Background()
		idHasherMock := mocks.NewMockIDHasher(control)
		urlRepoMock := mocks.NewMockUrlRepository(control)
		service := NewUrlService(urlRepoMock, idHasherMock)

		originalUrl := "http://example.com"
		id := 123
		expecterError := errors.New("failed to hash id")

		urlRepoMock.EXPECT().FindByOriginalUrl(ctx, originalUrl).Return(id, nil)
		idHasherMock.EXPECT().EncodeUrl(id).Return("", expecterError)

		shortCode, err := service.ShortenUrl(ctx, originalUrl)

		assert.Error(t, err)
		assert.Equal(t, expecterError, err)
		assert.Equal(t, "", shortCode)
	})

	t.Run("should return an error when repository saves an existing url, but hasher fails to encode id", func(t *testing.T) {
		control := gomock.NewController(t)
		defer control.Finish()

		ctx := context.Background()
		idHasherMock := mocks.NewMockIDHasher(control)
		urlRepoMock := mocks.NewMockUrlRepository(control)
		service := NewUrlService(urlRepoMock, idHasherMock)

		originalUrl := "http://example.com"
		id := 123
		expecterError := errors.New("failed to hash id")

		urlRepoMock.EXPECT().FindByOriginalUrl(ctx, originalUrl).Return(0, gorm.ErrRecordNotFound)
		urlRepoMock.EXPECT().Save(ctx, originalUrl).Return(id, nil)
		idHasherMock.EXPECT().EncodeUrl(id).Return("", expecterError)

		shortCode, err := service.ShortenUrl(ctx, originalUrl)

		assert.Error(t, err)
		assert.Equal(t, expecterError, err)
		assert.Equal(t, "", shortCode)
	})
}
