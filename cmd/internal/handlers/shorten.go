package handlers

import (
	"encoding/json"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/Rajit-Dutta/go-redis-url-shortener/cmd/api/helpers"
	"github.com/Rajit-Dutta/go-redis-url-shortener/cmd/internal/config"
	"github.com/Rajit-Dutta/go-redis-url-shortener/cmd/internal/db"
	"github.com/Rajit-Dutta/go-redis-url-shortener/cmd/internal/httpx"
	"github.com/redis/go-redis/v9"
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

	//implementation of rate limiting
	r2 := db.CreateClient(1)
	defer r2.Close()

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		host = r.RemoteAddr
	}

	val, err := r2.Get(db.Ctx, host).Result()
	if err == redis.Nil {
		err = r2.Set(db.Ctx, host, config.MustLoad().Quota, time.Second*60*30).Err()
		if err != nil {
			httpx.Error(w, http.StatusInternalServerError, "Something went wrong during IP fetching", "unsuccesful_fetch")
		}
	} else {
		valInt, err := strconv.Atoi(val)
		if err != nil {
			httpx.Error(w, http.StatusInternalServerError, "Something went wrong during host conversion to INT", "unsuccesful_conversion")
		}
		if valInt == 0 {
			ttl, _ := r2.TTL(db.Ctx, host).Result()
			httpx.RateLimitRestError(w, http.StatusBadRequest, "Rate limit exceeded", "rate_limit_exceeded", time.Duration(ttl.Seconds()))
		}
	}

	//check if actual URL

	if !helpers.RemoveDomainErrors(req.URL) {
		httpx.Error(w, http.StatusBadRequest, "URL is not valid", "invalid_URL")
	}

	req.URL = helpers.EnforceHTTP(req.URL)
}
