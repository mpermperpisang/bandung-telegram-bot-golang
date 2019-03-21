package database

import (
	"database/sql"

	"github.com/bandung-telegram-bot-golang/src/helper"
	_ "github.com/go-sql-driver/mysql"
)

func NewOnboarding(username string) bool {
	var onboarding bool
	var count int

	db := DBConnection()
	rowCount := db.QueryRow("SELECT on_flag from onboarding_member where on_username='@" + username + "'").Scan(&count)
	row, err := db.Query("SELECT on_flag from onboarding_member where on_username='@" + username + "'")
	helper.ErrorMessage(err)

	if rowCount == sql.ErrNoRows {
		return true
	} else {
		for row.Next() {
			err = row.Scan(&onboarding)
			helper.ErrorMessage(err)
		}
	}

	return onboarding
}
