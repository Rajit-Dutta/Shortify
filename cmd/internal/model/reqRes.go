package model

import "time"

type Request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type Response struct {
	URL             string        `json:"url"`
	CustomShort     string        `json:"short"`
	Expiry          time.Duration `json:"expiry"`
	XRateRemaining  int           `json:"rate_limit"`
	XRateLimitReset time.Duration `json:"rate_limit_rest"`
}

type ToBePublishedURL struct {
	OriginalURL  string        `json:"original_url"`
	ShortenedURL string        `json:"shortened_url"`
	Expiry       time.Duration `json:"expiry"`
}
