package db

import (
	"github.com/bandung-telegram-bot-golang/src/helper"
	_ "github.com/go-sql-driver/mysql"
)

func AddOnboarding(username string) {
	db := DBConnection()
	_, err := db.Exec("INSERT INTO onboarding_member (on_username, on_flag) VALUES ('" + username + "', 'false')")
	helper.ErrorMessage(err)
}
