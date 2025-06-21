package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kabakadev/goshort/internal/service"
)

type ShortenRequest struct {
	URL string `json:"url"`
}

func Shorten(c *fiber.Ctx) error {
	var body ShortenRequest

	if err := c.BodyParser(&body); err != nil || body.URL == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	code, err := service.ShortenURL(body.URL)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not shorten URL"})
	}

	return c.JSON(fiber.Map{
		"short_code": code,
		"short_url":  "http://localhost:3000/" + code,
	})
}

func Resolve(c *fiber.Ctx) error {
	code := c.Params("code")

	originalURL, err := service.ResolveURL(code)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "URL not found"})
	}

	return c.Redirect(originalURL, fiber.StatusFound)
}
