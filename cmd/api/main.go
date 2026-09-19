package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/Rajit-Dutta/go-redis-url-shortener/cmd/internal/config"
	"github.com/Rajit-Dutta/go-redis-url-shortener/cmd/internal/handlers"
)

func main() {
	//Setting up env
	cfg := config.MustLoad()

	//Route mux
	mux := http.NewServeMux()

	mux.HandleFunc("GET /:url", handlers.ResolveURL)
	mux.HandleFunc("POST /api/v1", handlers.ShortenURL)

	//Setting up the server
	server := http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	//Setting up the handler
	loggerHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	})

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server not running: %v", err)
	}

	logger := slog.New(loggerHandler)
	slog.SetDefault(logger)
}
