package service

import (
	"context"
	"errors"
	"testing"

	"github.com/Bromolima/url-shortner-go/internal/model"
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

func TestUrlService_Redirect(t *testing.T) {
	t.Run("should decode short code and return its original url", func(t *testing.T) {
		control := gomock.NewController(t)
		defer control.Finish()

		ctx := context.Background()
		idHasherMock := mocks.NewMockIDHasher(control)
		urlRepoMock := mocks.NewMockUrlRepository(control)
		service := NewUrlService(urlRepoMock, idHasherMock)

		shortCode := "abcde"
		id := 123
		expectedOriginalUrl := "http://example.com"

		idHasherMock.EXPECT().DecodeUrl(shortCode).Return(123, nil)
		urlRepoMock.EXPECT().Find(ctx, id).Return(expectedOriginalUrl, nil)

		originalUrl, err := service.Redirect(ctx, shortCode)

		assert.NoError(t, err)
		assert.Equal(t, originalUrl, expectedOriginalUrl)
	})

	t.Run("should return an error when hasher fails to decode url", func(t *testing.T) {
		control := gomock.NewController(t)
		defer control.Finish()

		ctx := context.Background()
		idHasherMock := mocks.NewMockIDHasher(control)
		urlRepoMock := mocks.NewMockUrlRepository(control)
		service := NewUrlService(urlRepoMock, idHasherMock)

		shortCode := "abcde"
		expectedError := errors.New("failed to decode code")

		idHasherMock.EXPECT().DecodeUrl(shortCode).Return(0, expectedError)

		originalUrl, err := service.Redirect(ctx, shortCode)

		assert.Equal(t, err, expectedError)
		assert.Equal(t, originalUrl, "")
	})

	t.Run("shoud return url not doun when url is not in database", func(t *testing.T) {
		control := gomock.NewController(t)
		defer control.Finish()

		ctx := context.Background()
		idHasherMock := mocks.NewMockIDHasher(control)
		urlRepoMock := mocks.NewMockUrlRepository(control)
		service := NewUrlService(urlRepoMock, idHasherMock)

		shortCode := "abcde"
		id := 123
		expectedError := model.ErrUrlNotFound

		idHasherMock.EXPECT().DecodeUrl(shortCode).Return(id, nil)
		urlRepoMock.EXPECT().Find(ctx, id).Return("", gorm.ErrRecordNotFound)

		originalUrl, err := service.Redirect(ctx, shortCode)

		assert.Equal(t, err, expectedError)
		assert.Equal(t, originalUrl, "")
	})

	t.Run("shoud return an error when repository fails to find original url", func(t *testing.T) {
		control := gomock.NewController(t)
		defer control.Finish()

		ctx := context.Background()
		idHasherMock := mocks.NewMockIDHasher(control)
		urlRepoMock := mocks.NewMockUrlRepository(control)
		service := NewUrlService(urlRepoMock, idHasherMock)

		shortCode := "abcde"
		id := 123
		expectedError := errors.New("database connection failed")

		idHasherMock.EXPECT().DecodeUrl(shortCode).Return(id, nil)
		urlRepoMock.EXPECT().Find(ctx, id).Return("", expectedError)

		originalUrl, err := service.Redirect(ctx, shortCode)

		assert.Equal(t, err, expectedError)
		assert.Equal(t, originalUrl, "")
	})
}
