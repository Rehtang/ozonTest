package utils

import (
	"net/url"
)

func ValidateURL(rawURL string) bool {
	_, err := url.ParseRequestURI(rawURL)
	return err == nil
}
