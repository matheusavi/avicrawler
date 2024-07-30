package main

import (
	"log"

	"github.com/avicrawler/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/crawl", handlers.HandleCrawUrl)

	log.Fatal(app.Listen(":3457"))
}
