package dto

type ShortenUrlPayload struct {
	OriginalUrl string `json:"url"`
}

type UrlResponse struct {
	ShortCode string `json:"short_code"`
}

type RedirectUrlResponse struct {
	ShortCode string `uri:"short_code"`
}
