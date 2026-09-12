package handlers

import (
	"net/http"
	"time"
)

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type response struct {
	URL             string        `json:"url"`
	CustomShort     string        `json:"short"`
	Expiry          time.Duration `json:"expiry"`
	XRateRemaining  int           `json:"rate_limit"`
	XRateLimitReset time.Duration `json:"rate_limit_rest"`
}

func ResolveURL(w http.ResponseWriter, r *http.Request) {

}

func ShortenURL(w http.ResponseWriter, r *http.Request) {

}
