package httpx

import (
	"encoding/json"
	"net/http"
)

type ErrorPayload struct {
	Field   string `json:"field,omitempty"`
	Code    string `json:"code"`
	Message string `json:"msg"`
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
