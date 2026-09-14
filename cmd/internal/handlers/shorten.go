package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Rajit-Dutta/go-redis-url-shortener/cmd/api/helpers"
	"github.com/Rajit-Dutta/go-redis-url-shortener/cmd/internal/httpx"
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

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	var req request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.Error(w, http.StatusBadRequest, "Something went wrong while reading body", "invalid_body")
	}

	//check if actual URL

	if !helpers.RemoveDomainErrors(req.URL) {
		httpx.Error(w, http.StatusBadRequest, "URL is not valid", "invalid_URL")
	}

	req.URL = helpers.EnforceHTTP(req.URL)
}
