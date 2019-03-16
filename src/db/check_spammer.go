package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func CheckSpammer(username, command string) bool {
	var count int

	db := DBConnection()

	row := db.QueryRow("SELECT * FROM bot_spam WHERE spammer_name='" + username + "' and bot_command='" + command + "' and bot_attempt>5").Scan(&count)

	if row == sql.ErrNoRows {
		return false
	} else {
		return true
	}
}
