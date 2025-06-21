package service

import (
	"github.com/google/uuid"
	"github.com/kabakadev/goshort/internal/repository"
)

func ShortenURL(originalURL string) (string, error) {
	shortCode := uuid.New().String()[:6] // generate 6-char short code
	err := repository.SaveURL(shortCode, originalURL)
	if err != nil {
		return "", err
	}
	return shortCode, nil
}

func ResolveURL(code string) (string, error) {
	return repository.GetOriginalURL(code)
}
