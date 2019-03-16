package user

import (
	"github.com/bandung-telegram-bot-golang/src/db"
)

func IsAdmin(username string) bool {
	query := db.CheckAdmin(username)

	return query
}

func IsSpammer(username, command string) bool {
	query := db.CheckSpammer(username, command)

	return query
}

func SaveSpammer(username, command string) {
	db.SaveSpammer(username, command)
}
