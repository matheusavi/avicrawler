package handlers

import (
	"strconv"

	"github.com/avicrawler/pkg"
	"github.com/gofiber/fiber/v2"
)

func HandleCrawUrl(c *fiber.Ctx) error {
	url := c.Query("url", "http://localhost:3000/")
	i := pkg.CrawlUrl(url, 2, c.Context())
	return c.SendString(strconv.Itoa(i) + " urls crawled under " + url)
}
