package database

import (
	"github.com/bandung-telegram-bot-golang/src/helper"
	_ "github.com/go-sql-driver/mysql"
)

func DeleteOnboarding(username string) {
	db := DBConnection()
	_, err := db.Exec("DELETE FROM onboarding_member WHERE on_username='" + username + "'")
	helper.ErrorMessage(err)
}
