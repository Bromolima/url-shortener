package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Bromolima/url-shortner-go/internal/http/dto"
	"github.com/Bromolima/url-shortner-go/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter(control *gomock.Controller) (*gin.Engine, *mocks.MockUrlService) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	urlServiceMock := mocks.NewMockUrlService(control)
	urlHandler := NewUrlHandler(urlServiceMock)

	router.POST("/v1/shorten", urlHandler.ShortenUrl)

	return router, urlServiceMock
}

func TestUrlHandler_ShortenUrl(t *testing.T) {
	t.Run("should return an status created and the shortned url", func(t *testing.T) {
		control := gomock.NewController(t)
		defer control.Finish()

		router, urlServiceMock := setupTestRouter(control)

		expectedPayload := dto.ShortenUrlPayload{
			OriginalUrl: "http://example.com",
		}
		urlBody, _ := json.Marshal(expectedPayload)

		expectedShortUrl := "http://short.ly/abc123"

		urlServiceMock.EXPECT().
			ShortenUrl(gomock.Any(), expectedPayload.OriginalUrl).
			Return(expectedShortUrl, nil)

		req, _ := http.NewRequest(http.MethodPost, "/v1/shorten", bytes.NewReader(urlBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		var actualResponse dto.UrlResponse
		if err := json.NewDecoder(w.Body).Decode(&actualResponse); err != nil {
			t.Fatalf("erro ao decodificar resposta: %v", err)
		}

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Equal(t, expectedShortUrl, actualResponse.ShortCode)
	})

	t.Run("should return status bad request when body is invalid", func(t *testing.T) {
		control := gomock.NewController(t)
		defer control.Finish()

		router, _ := setupTestRouter(control)

		invalidPayload := []byte(
			`"originalUrl": "http://example.com,"`,
		)

		req, _ := http.NewRequest(http.MethodPost, "/v1/shorten", bytes.NewReader(invalidPayload))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})

	t.Run("should return status unprocessable entity when body is invalid", func(t *testing.T) {
		control := gomock.NewController(t)
		defer control.Finish()

		router, _ := setupTestRouter(control)

		invalidPayload := []byte(`"url": "http://example.com"`)

		req, _ := http.NewRequest(http.MethodPost, "/v1/shorten", bytes.NewReader(invalidPayload))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})

	t.Run("should return status bad request when url format is invalid", func(t *testing.T) {
		control := gomock.NewController(t)
		defer control.Finish()

		router, _ := setupTestRouter(control)

		invalidUrl := []byte(
			`{"url": "invalidUrl"}`,
		)

		req, _ := http.NewRequest(http.MethodPost, "/v1/shorten", bytes.NewReader(invalidUrl))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("should return an internal server error when service fails to save url", func(t *testing.T) {
		control := gomock.NewController(t)
		defer control.Finish()

		router, urlServiceMock := setupTestRouter(control)

		expectedPayload := dto.ShortenUrlPayload{
			OriginalUrl: "http://example.com",
		}
		urlBody, _ := json.Marshal(expectedPayload)

		internalError := errors.New("internal server error")

		urlServiceMock.EXPECT().
			ShortenUrl(gomock.Any(), expectedPayload.OriginalUrl).
			Return("", internalError)

		req, _ := http.NewRequest(http.MethodPost, "/v1/shorten", bytes.NewReader(urlBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
