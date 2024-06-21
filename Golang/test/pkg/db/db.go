package db

import (
	"database/sql"
	"log"
	"test/internal/config"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"

	_ "github.com/lib/pq"
)

var DB *reform.DB

func InitDB() {
	db, err := sql.Open("postgres", config.AppConfig.DBSource)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	DB = reform.NewDB(db, postgresql.Dialect, nil)
}
