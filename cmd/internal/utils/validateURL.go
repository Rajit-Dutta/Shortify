package utils

import (
	"errors"
	"net/url"
)

func ValidateURL(reqURL string) error {
	parsedURL, err := url.ParseRequestURI(reqURL)

	if err != nil {
		return errors.New("invalid URL")
	}

	if parsedURL.Scheme != "http" &&
		parsedURL.Scheme != "https" {
		return errors.New("URL must use HTTP or HTTPS")
	}

	if parsedURL.Host == "" {
		return errors.New("URL must contain a host")
	}

	return nil
}
