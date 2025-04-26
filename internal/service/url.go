package service

import (
	"errors"
)

var urlStore = make(map[string]string)

func SetURL(shortcode, originalURL string) {
	urlStore[shortcode] = originalURL
}

func GetURL(shortcode string) (string, error) {
	originalURL, ok := urlStore[shortcode]
	if !ok {
		return "", errors.New("shortcode not found")
	}

	return originalURL, nil
}
