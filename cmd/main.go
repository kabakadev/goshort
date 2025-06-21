package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kabakadev/goshort/config"
	"github.com/kabakadev/goshort/internal/handler"


)

func main() {
	// Step 1: Create a new Fiber app
	app := fiber.New()

	config.InitDB()


	// Step 2: Define a health check route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("GoShort is running 🚀")
	})
	app.Post("/shorten", handler.Shorten)
	app.Get("/:code", handler.Resolve)


	// Step 3: Start the server on port 3000
	log.Println("Server is running at http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
