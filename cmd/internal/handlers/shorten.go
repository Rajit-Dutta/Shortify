package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/Rajit-Dutta/go-redis-url-shortener/cmd/internal/httpx"
	"github.com/Rajit-Dutta/go-redis-url-shortener/cmd/internal/middleware"
	"github.com/Rajit-Dutta/go-redis-url-shortener/cmd/internal/model"
	"github.com/Rajit-Dutta/go-redis-url-shortener/cmd/internal/repository"
	"github.com/Rajit-Dutta/go-redis-url-shortener/cmd/internal/service"
)

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	var req model.Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.Error(w, http.StatusBadRequest, "Something went wrong while reading body", "invalid_body")
		return
	}

	result, err := service.PublishURL(req, r.Context())
	if err != nil {
		if errors.Is(err, repository.ErrCustomShortExists) {
			httpx.Error(w, http.StatusConflict, err.Error(), "url_short_in_use")
			return
		}
		httpx.Error(w, http.StatusBadRequest, err.Error(), "invalid_url")
		return
	}

	resp := model.Response{
		URL:             result.OriginalURL,
		CustomShort:     result.ShortenedURL,
		Expiry:          result.Expiry,
		XRateRemaining:  int(middleware.RateLimitRemaining(r.Context())),
		XRateLimitReset: time.Duration(middleware.RateLimitReset(r.Context())),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
