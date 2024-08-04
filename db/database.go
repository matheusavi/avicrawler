package db

import (
	"context"
	"database/sql"

	"github.com/gofiber/fiber/v2/log"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var store *bun.DB

func InitializeStore(dsn string) error {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	err := sqldb.Ping()
	if err != nil {
		return err
	}

	store = bun.NewDB(sqldb, pgdialect.New())

	return nil
}

type Site struct {
	ID      int32 `bun:"id,pk,autoincrement"`
	Url     string
	Content string
}

func SaveContent(s *Site, c context.Context) {
	_, err := store.NewInsert().Model(s).Exec(c)

	if err != nil {
		log.Error(err)
	}
}
