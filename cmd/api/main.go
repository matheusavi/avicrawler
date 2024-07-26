package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/crawl", func(c *fiber.Ctx) error {
		return c.SendString("Crawling " + c.Query("url", "http://localhost:3000/"))
	})

	log.Fatal(app.Listen(":3000"))
}
