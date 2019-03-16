package db

import (
	"database/sql"
	"os"

	"github.com/bandung-telegram-bot-golang/src/helper"
	_ "github.com/go-sql-driver/mysql"
)

func DBConnection() (db *sql.DB) {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	schema := os.Getenv("DB_SCHEMA")
	db, err := sql.Open("mysql", user+":"+password+"@/"+schema)
	helper.ErrorMessage(err)

	return db
}
