package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("usage: go run ./cmd/migrate <command> [args]")
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is empty (example: postgres://shop_user:shop_pass@localhost:5432/shopdb?sslmode=disable)")
	}

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("cannot connect to DB: ", err)
	}

	// migrations лежат в папке migrations
	dir := "migrations"

	// передаём все аргументы (кроме имени программы) в goose
	if err := goose.Run(os.Args[1], db, dir, os.Args[2:]...); err != nil {
		log.Fatal(err)
	}
}
