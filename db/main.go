package db

import (
	"database/sql"
	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"log"
	"os"
)

var DB *bun.DB

func InitDB() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dsn := "postgres://" + os.Getenv("DB_USER") + ":" +
		os.Getenv("DB_PASSWORD") + "@" +
		os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + "/" +
		os.Getenv("DB_NAME") + "?sslmode=disable"

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	DB = bun.NewDB(sqldb, pgdialect.New())
	log.Println("Database connected successfully!")
}
