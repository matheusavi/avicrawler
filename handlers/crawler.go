package handlers

import (
	"github.com/avicrawler/pkg"
	"github.com/gofiber/fiber/v2"
)

func HandleCrawUrl(c *fiber.Ctx) error {
	url := c.Query("url", "http://localhost:3000/")
	pkg.CrawlUrl(url, 2)
	return c.SendString("Crawling " + url)
}
