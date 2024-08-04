package main

import (
	"log"

	"github.com/avicrawler/db"
	"github.com/avicrawler/handlers"
	"github.com/avicrawler/types"
	"github.com/gofiber/fiber/v2"
)

func main() {
	var cfg types.InitialConfig
	err := cfg.ParseFromFile()

	if err != nil {
		log.Fatal(err)
	}

	err = cfg.ParseFromEnv()

	if err != nil {
		log.Fatal(err)
	}

	err = db.InitializeStore(cfg.Database.Dsn)

	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/crawl", handlers.HandleCrawUrl)

	log.Fatal(app.Listen(cfg.Server.Host + ":" + cfg.Server.Port))
}
