package httpx

import (
	"encoding/json"
	"net/http"
	"time"
)

type ErrorPayload struct {
	Field         string        `json:"field,omitempty"`
	Code          string        `json:"code"`
	Message       string        `json:"msg"`
	RateLimitRest time.Duration `json:"rate_limit_rest,omitempty"`
}

type ErrorWrapper struct {
	Error ErrorPayload `json:"error"`
}

func Error(w http.ResponseWriter, status int, message string, code string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(ErrorWrapper{
		ErrorPayload{
			Code:    code,
			Message: message,
		},
	})
}

func RateLimitRestError(w http.ResponseWriter, status int, message string, code string, rateLimit time.Duration) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(ErrorWrapper{
		ErrorPayload{
			Code:          code,
			Message:       message,
			RateLimitRest: rateLimit,
		},
	})
}
