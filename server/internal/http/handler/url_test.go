package handler

import (
	"context"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Bromolima/url-shortner-go/internal/mocks"
	"github.com/Bromolima/url-shortner-go/internal/model"
	"github.com/golang/mock/gomock"
)

func Test_ShortenUrl(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()

	serviceMock := mocks.NewMockUrlService(ctrl)
	urlHandler := NewUrlHandler(serviceMock)

	testCases := []struct {
		name           string
		body           string
		setupMocks     func()
		expectedStatus int
	}{
		{
			name: "when there is no error",
			body: `{"url": "https://example.com"}`,
			setupMocks: func() {
				serviceMock.EXPECT().ShortenUrl(ctx, "https://example.com").Return("abcdef", nil)
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name:           "when there is an error to parse the requisition body",
			body:           `{"example"}`,
			setupMocks:     func() {},
			expectedStatus: http.StatusUnprocessableEntity,
		},
		{
			name:           "when the url is not valid",
			body:           `{"url": ""}`,
			setupMocks:     func() {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "when there is an internal error",
			body: `{"url": "https://example.com"}`,
			setupMocks: func() {
				serviceMock.EXPECT().ShortenUrl(ctx, "https://example.com").Return("abcdef", sql.ErrConnDone)
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMocks()

			rec := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/shorten", strings.NewReader(tt.body))
			req.Header.Set("Content-Type", "application/json")
			req.Host = "localhost:8080"

			urlHandler.ShortenUrl(rec, req)

			if rec.Code != tt.expectedStatus {
				t.Error(rec.Code)
			}
		})
	}
}

func Test_Redirect(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()

	serviceMock := mocks.NewMockUrlService(ctrl)
	urlHandler := NewUrlHandler(serviceMock)

	testCases := []struct {
		name           string
		shortCode      string
		originalUrl    string
		setupMocks     func()
		expectedStatus int
	}{
		{
			name:        "when there is no error",
			shortCode:   "abcdef",
			originalUrl: "/originalUrl",
			setupMocks: func() {
				serviceMock.EXPECT().Redirect(ctx, "abcdef").Return("/originalUrl", nil)
			},
			expectedStatus: http.StatusFound,
		},
		{
			name:        "when the original url is not found",
			shortCode:   "abcdef",
			originalUrl: "",
			setupMocks: func() {
				serviceMock.EXPECT().Redirect(ctx, "abcdef").Return("", model.ErrUrlNotFound)
			},
			expectedStatus: http.StatusNotFound,
		},
		{
			name:        "when there is an internal error",
			shortCode:   "abcdef",
			originalUrl: "",
			setupMocks: func() {
				serviceMock.EXPECT().Redirect(ctx, "abcdef").Return("", sql.ErrConnDone)
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range testCases {
		tt.setupMocks()

		rec := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/"+tt.shortCode, nil)

		urlHandler.Redirect(rec, req)
		if rec.Code != tt.expectedStatus {
			t.Error(rec.Code)
		}
	}
}
