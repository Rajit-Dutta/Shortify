package repository

import (
	"context"
	"errors"
	"time"

	"github.com/Rajit-Dutta/go-redis-url-shortener/cmd/internal/db"
	"github.com/redis/go-redis/v9"
)

var ErrCustomShortExists = errors.New("custom short URL is already in use")

func IsExistCustomShortURL(urlID string) error {

	conn := db.CreateClient(0)
	defer conn.Close()

	_, err := conn.Get(db.Ctx, urlID).Result()

	if err == redis.Nil {
		return nil
	}

	if err != nil {
		return err
	}
	return ErrCustomShortExists

}

func SetShortenedURL(ctx context.Context, urlID string, url string, urlExpiry time.Duration) error {
	conn := db.CreateClient(0)
	defer conn.Close()

	if err := conn.Set(ctx, urlID, url, urlExpiry*3600*time.Second).Err(); err != nil {
		return err
	}
	return nil
}
