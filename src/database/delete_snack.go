package database

import (
	"database/sql"
	"regexp"

	"github.com/bandung-telegram-bot-golang/src/helper"
	_ "github.com/go-sql-driver/mysql"
)

func DeleteSnack(snack string) {
	var snackArray []string
	var count int

	db := DBConnection()
	file := helper.CreateFile()
	listUsername := regexp.MustCompile(helper.RegexCompileUsername()).FindAllString(snack, -1)

	for _, username := range listUsername {
		includeSnack, _ := helper.IncludeArray(username, snackArray)

		if !includeSnack {
			row := db.QueryRow("SELECT * FROM bandung_snack WHERE name='" + username + "'").Scan(&count)

			if row != sql.ErrNoRows {
				_, err := db.Exec("DELETE FROM bandung_snack WHERE name='" + username + "'")
				helper.ErrorMessage(err)

				file.WriteString("\n- ByBy Kak " + username)
			} else {
				file.WriteString("\n- <code>" + username + "</code> siapa tuh? 👻")
			}

			snackArray = append(snackArray, username)
		}
	}

	defer file.Close()
}
