package dto

type UrlPayload struct {
	OriginalUrl string `json:"url"`
}

type UrlResponse struct {
	ShortCode string `json:"short_code"`
}
