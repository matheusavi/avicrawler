package main

import (
	"log"

	"github.com/avicrawler/handlers"
	"github.com/avicrawler/types"
	"github.com/gofiber/fiber/v2"
)

func main() {
	var cfg types.InitialConfig
	cfg.ParseFromFile()
	cfg.ParseFromEnv()

	app := fiber.New()

	app.Get("/crawl", handlers.HandleCrawUrl)

	log.Fatal(app.Listen(cfg.Server.Host + ":" + cfg.Server.Port))
}
