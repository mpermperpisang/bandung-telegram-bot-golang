package user

import (
	"db"
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
