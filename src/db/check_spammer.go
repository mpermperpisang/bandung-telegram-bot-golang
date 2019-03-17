package db

import (
	"database/sql"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func CheckSpammer(username, command string) bool {
	var count, attempt int

	db := DBConnection()
	admin := CheckAdmin(username)

	if admin {
		attempt = 25
	} else {
		attempt = 5
	}

	row := db.QueryRow("SELECT * FROM bot_spam WHERE spammer_name='" + username + "' and bot_command='" + command + "' and bot_attempt>" + strconv.Itoa(attempt) + "").Scan(&count)

	if row == sql.ErrNoRows {
		return false
	} else {
		return true
	}
}
