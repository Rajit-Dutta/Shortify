package helpers

import (
	"strings"

	"github.com/Rajit-Dutta/go-redis-url-shortener/cmd/internal/config"
)

func EnforceHTTP(url string) string {
	if url[:4] != "http" {
		return "http://" + url
	}
	return url
}

func RemoveDomainErrors(url string) bool {
	if url == config.MustLoad().Domain {
		return false
	}

	newURL := strings.Replace(url, "http://", "", 1)
	newURL = strings.Replace(newURL, "https://", "", 1)
	newURL = strings.Replace(newURL, "www.", "", 1)
	newURL = strings.Split(newURL, "/")[0]

	if newURL == config.MustLoad().Domain {
		return false
	}

	return true
}
