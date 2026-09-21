package utils

import "github.com/google/uuid"

func GenerateCustomShort(url string) string {
	id := uuid.New().String()[:4]
	return id
}
