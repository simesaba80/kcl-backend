package main

import (
	"database/sql"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/simesaba80/kcl-back/services"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

func connectdb() {
	// Bunを使ってDB接続
	dsn := "postgres://postgres:postgres@db:5432/postgres?sslmode=disable"
	sqlDB := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqlDB, pgdialect.New())
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return
	}

	// クエリーフックを追加
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
	))
}

func main() {
	e := echo.New()
	connectdb()

	e.GET("/", services.Hello)

	e.Start(":8080")
}
