package database

import "github.com/bandung-telegram-bot-golang/src/helper"

func UpdateOnboarding(username string) {
	db := DBConnection()
	_, err := db.Exec("UPDATE onboarding_member SET on_flag='true' WHERE on_username='@" + username + "'")
	helper.ErrorMessage(err)
}
