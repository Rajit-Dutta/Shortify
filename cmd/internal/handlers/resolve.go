package handlers

import (
	"net/http"

	"github.com/Rajit-Dutta/go-redis-url-shortener/cmd/internal/db"
	"github.com/Rajit-Dutta/go-redis-url-shortener/cmd/internal/httpx"
	"github.com/redis/go-redis/v9"
)

func ResolveURL(w http.ResponseWriter, r *http.Request) {
	url := r.PathValue("url")
	rd := db.CreateClient(0)
	defer rd.Close()

	value, err := rd.Get(db.Ctx, url).Result()
	if err == redis.Nil {
		httpx.Error(w, http.StatusNotFound, "shortened URL not found", "not_found")
	} else if err != nil {
		httpx.Error(w, http.StatusInternalServerError, "cannot connect to DB", "connection_error")
	}

	rInr := db.CreateClient(1)
	defer rInr.Close()

	_ = rInr.Incr(db.Ctx, "counter")

	http.Redirect(w, r, value, http.StatusMovedPermanently)
}
