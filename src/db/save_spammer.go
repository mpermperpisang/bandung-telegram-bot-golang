package db

import (
	"database/sql"
	"helper"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func SaveSpammer(username, command string) {
	var count, attempt, sumAttempt int

	db := DBConnection()

	rowCommand := db.QueryRow("SELECT * FROM bot_spam WHERE bot_command='" + command + "'").Scan(&count)
	rowUsername := db.QueryRow("SELECT * FROM bot_spam WHERE spammer_name='" + username + "' and bot_command='" + command + "'").Scan(&count)
	countAttempt, err := db.Query("SELECT bot_attempt FROM bot_spam WHERE bot_command='" + command + "'")
	helper.ErrorMessage(err)

	for countAttempt.Next() {
		err = countAttempt.Scan(&attempt)
		helper.ErrorMessage(err)

		sumAttempt = attempt + 1
	}

	if rowCommand == sql.ErrNoRows {
		db.Exec("INSERT INTO bot_spam VALUES (1, '" + username + "', '" + command + "')")
	} else if rowUsername == sql.ErrNoRows {
		db.Exec("UPDATE bot_spam SET bot_attempt=1, spammer_name='" + username + "'  WHERE bot_command='" + command + "'")
	} else {
		db.Exec("UPDATE bot_spam SET bot_attempt=" + strconv.Itoa(sumAttempt) + " WHERE spammer_name='" + username + "'  and bot_command='" + command + "'")
	}
}
