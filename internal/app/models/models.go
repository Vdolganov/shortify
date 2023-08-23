package models

type ShorterRequestBody struct {
	URL string `json:"url"`
}

type ShorterResponse struct {
	Result string `json:"result"`
}
