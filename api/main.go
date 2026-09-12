package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/Rajit-Dutta/go-redis-url-shortener/internal/handlers"
)

func main() {
	//Setting up the server
	server := http.Server{
		Addr: ":8000",
		//Handler: handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server not running: %v", err)
	}

	//Route mux
	mux := http.NewServeMux()

	mux.HandleFunc("GET /:url", handlers.ResolveURL)
	mux.HandleFunc("POST /api/v1", handlers.ShortenURL)

	//Setting up the handler
	loggerHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	})

	logger := slog.New(loggerHandler)
	slog.SetDefault(logger)
}
