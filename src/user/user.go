package user

import "github.com/bandung-telegram-bot-golang/src/database"

func IsAdmin(username string) bool {
	query := database.CheckAdmin(username)

	return query
}

func IsSpammer(username, command string) bool {
	query := database.CheckSpammer(username, command)

	return query
}

func SaveSpammer(username, command string) {
	database.SaveSpammer(username, command)
}
