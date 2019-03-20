package database

import (
	"github.com/bandung-telegram-bot-golang/src/helper"
	_ "github.com/go-sql-driver/mysql"
)

func NewOnboarding(username string) bool {
	var onboarding bool

	db := DBConnection()

	row, err := db.Query("SELECT on_flag from onboarding_member where on_username='@" + username + "'")
	helper.ErrorMessage(err)

	for row.Next() {
		err = row.Scan(&onboarding)
		helper.ErrorMessage(err)
	}

	return onboarding
}
