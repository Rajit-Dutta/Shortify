package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type config struct {
	Port   string
	Domain string
	//Env         string
	Quota       int
	DatabaseURL string
}

func MustLoad() config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("godotenv.load: Could not load env")
	}
	port := os.Getenv("APP_PORT")
	if port == "" {
		panic("PORT is required")
	}
	domain := os.Getenv("DOMAIN")
	if domain == "" {
		panic("DOMAIN is required")
	}
	db := os.Getenv("DB_ADDR")
	if db == "" {
		panic("DB_ADDR is required")
	}
	quota := os.Getenv("APP_QUOTA")
	if quota == "" {
		panic("QUOTA is required")
	}
	quotaValue, err := strconv.Atoi(quota)
	if err != nil {
		panic("QUOTA must be an integer")
	}

	return config{
		Port:        port,
		Domain:      domain,
		DatabaseURL: db,
		Quota:       quotaValue,
	}
}
