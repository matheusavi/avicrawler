package db

import (
	"os"

	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
)

func SaveContent(url string, content string) {
	err := os.WriteFile("./html/"+uuid.New().String()+".html", []byte(content), 0644)
	log.Error(err)
}
