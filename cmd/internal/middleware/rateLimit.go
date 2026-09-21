package middleware

import (
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

		val, err := conn.Get(db.Ctx, host).Result()
		if err == redis.Nil {
			err = conn.Set(db.Ctx, host, config.MustLoad().Quota, time.Second*60*30).Err()
			if err != nil {
				httpx.Error(w, http.StatusInternalServerError, "Something went wrong during IP fetching", "unsuccesful_fetch")
				return
			}
		} else {
			valInt, err := strconv.Atoi(val)
			if err != nil {
				httpx.Error(w, http.StatusInternalServerError, "Something went wrong during host conversion to INT", "unsuccesful_conversion")
				return
			}
			if valInt == 0 {
				ttl, _ := conn.TTL(db.Ctx, host).Result()
				httpx.RateLimitRestError(w, http.StatusBadRequest, "Rate limit exceeded", "rate_limit_exceeded", time.Duration(ttl.Seconds()))
				return
			}
		}
	})
}
