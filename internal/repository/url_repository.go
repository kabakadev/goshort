package repository

import (
	"database/sql"
	"github.com/kabakadev/goshort/config"
	"github.com/kabakadev/goshort/internal/model"
)

func SaveURL(shortCode string, originalURL string) error {
	query := `INSERT INTO urls (short_code, original_url) VALUES (?, ?)`

	_, err := config.DB.Exec(query, shortCode, originalURL)
	return err
}

func GetOriginalURL(shortCode string) (string, error) {
	var originalURL string
	query := `SELECT original_url FROM urls WHERE short_code = ?`

	err := config.DB.QueryRow(query, shortCode).Scan(&originalURL)
	if err != nil {
		return "", err
	}

	return originalURL, nil
}
