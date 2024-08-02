package main

import (
	"log"
	"os"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	dir, err := os.Getwd()
	dir = strings.ReplaceAll(dir, "\\", "/")
	log.Println("migrating")
	if err != nil {
		log.Fatal(err)
		return
	}
	migrationPath := "file:" + dir + "/cmd/migrate/migrations"
	log.Println(migrationPath)
	m, err := migrate.New(
		migrationPath,
		"postgres://postgres:avi123@db:5432/crawler?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}
