package middleware

import (
	"context"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/Rajit-Dutta/go-redis-url-shortener/cmd/internal/config"
	"github.com/Rajit-Dutta/go-redis-url-shortener/cmd/internal/db"
	"github.com/Rajit-Dutta/go-redis-url-shortener/cmd/internal/httpx"
	"github.com/redis/go-redis/v9"
)

func RateLimit(next http.Handler) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		conn := db.CreateClient(1)
		defer conn.Close()

		host, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			host = r.RemoteAddr
		}

		val, err := conn.Get(r.Context(), host).Result()
		if err == redis.Nil {
			err = conn.Set(r.Context(), host, config.MustLoad().Quota, time.Second*60*30).Err()
			if err != nil {
				httpx.Error(w, http.StatusInternalServerError, "Something went wrong during setting quota for IP", "unsuccessful_fetch")
				return
			}
		} else if err != nil {
			httpx.Error(w, http.StatusInternalServerError, "Something went wrong during IP fetching", "unsuccessful_fetch")
			return
		} else {
			valInt, err := strconv.Atoi(val)
			if err != nil {
				httpx.Error(w, http.StatusInternalServerError, "Something went wrong during host conversion to INT", "unsuccessful_conversion")
				return
			}
			if valInt == 0 {
				ttl, err := conn.TTL(r.Context(), host).Result()
				if err != nil {
					httpx.Error(w, http.StatusInternalServerError, "Unable to read rate limit TTL", "rate_limit_error")
					return
				}
				httpx.RateLimitRestError(w, http.StatusTooManyRequests, "Rate limit exceeded", "rate_limit_exceeded", time.Duration(ttl.Seconds()))
				return
			}
		}

		remaining, err := conn.Decr(r.Context(), host).Result()
		if err != nil {
			httpx.Error(w, http.StatusInternalServerError, "Unable to update rate limit", "rate_limit_error")
			return
		}

		ttl, err := conn.TTL(r.Context(), host).Result()
		if err != nil {
			httpx.Error(w, http.StatusInternalServerError, "Unable to read rate limit TTL", "rate_limit_error")
			return
		}

		ctx := context.WithValue(r.Context(), "rate_limit", int(remaining))
		ctx = context.WithValue(ctx, "rate_limit_rest", int64(ttl.Seconds()))

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RateLimitRemaining(ctx context.Context) int {
	value, ok := ctx.Value("rate_limit").(int)
	if !ok {
		return 0
	}
	return value
}

func RateLimitReset(ctx context.Context) int64 {
	value, ok := ctx.Value("rate_limit_rest").(int64)
	if !ok {
		return 0
	}
	return value
}
