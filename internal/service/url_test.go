package service

import (
	"context"
	"database/sql"
	"testing"

	"github.com/Bromolima/url-shortner-go/internal/mocks"
	"go.uber.org/mock/gomock"
)

func Test_ShotenUrl(t *testing.T) {
	control := gomock.NewController(t)
	ctx := context.Background()

	repositoryMock := mocks.NewMockUrlRepository(control)
	service := NewUrlService(repositoryMock)

	testCases := []struct {
		name        string
		originalUrl string
		shortCode   string
		setupMock   func()
		expectErr   bool
	}{
		{
			name:        "quando não houver nenhum erro e o shortCode não for encontrado",
			originalUrl: "https://example.com",
			setupMock: func() {
				repositoryMock.EXPECT().FindByOriginalUrl(ctx, "https://example.com").Return("", sql.ErrNoRows)
				repositoryMock.EXPECT().Save(ctx, gomock.Any()).Return(nil)
			},
			expectErr: false,
		},
		{
			name:        "quando não houver nenhum erro e o shortCode for encontrado",
			originalUrl: "https://example.com",
			shortCode:   "abc1234",
			setupMock: func() {
				repositoryMock.EXPECT().FindByOriginalUrl(ctx, "https://example.com").Return("abc1234", nil)
			},
			expectErr: false,
		},
		{
			name:        "quando houver um erro inesperado ao buscar o shortCode",
			originalUrl: "https://example.com",
			setupMock: func() {
				repositoryMock.EXPECT().FindByOriginalUrl(ctx, "https://example.com").Return("", sql.ErrConnDone)
			},
			expectErr: true,
		},
		{
			name:        "quando houver um erro inesperado ao salvar a URL",
			originalUrl: "https://example.com",
			setupMock: func() {
				repositoryMock.EXPECT().FindByOriginalUrl(ctx, "https://example.com").Return("", sql.ErrNoRows)
				repositoryMock.EXPECT().Save(ctx, gomock.Any()).Return(sql.ErrConnDone)
			},
			expectErr: true,
		},
	}

	for _, tt := range testCases {
		ctx := context.Background()
		tt.setupMock()

		_, err := service.ShortenUrl(ctx, tt.originalUrl)
		if err != nil && !tt.expectErr {
			t.Error(err)
		}
	}

}

func Test_Redirect(t *testing.T) {
	control := gomock.NewController(t)
	ctx := context.Background()

	repositoryMock := mocks.NewMockUrlRepository(control)
	service := NewUrlService(repositoryMock)

	testCases := []struct {
		name      string
		shortCode string
		setupMock func()
		expectErr bool
	}{
		{
			name:      "quando não houver nenhum erro",
			shortCode: "abc1234",
			setupMock: func() {
				repositoryMock.EXPECT().FindByShortCode(ctx, "abc1234").Return("https://example.com", nil)
			},
			expectErr: false,
		},
		{
			name:      "quando houver um erro inesperado ao buscar a URL",
			shortCode: "abc1234",
			setupMock: func() {
				repositoryMock.EXPECT().FindByShortCode(ctx, "abc1234").Return("", sql.ErrConnDone)
			},
			expectErr: true,
		},
	}

	for _, tt := range testCases {
		ctx := context.Background()
		tt.setupMock()

		_, err := service.Redirect(ctx, tt.shortCode)
		if err != nil && !tt.expectErr {
			t.Error(err)
		}
	}
}
