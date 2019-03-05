package db

import (
	"database/sql"
	"helper"
	"regexp"

	_ "github.com/go-sql-driver/mysql"
)

func CancelSnack(snack string) {
	var snackArray []string
	var count int

	db := DBConnection()
	file := helper.CreateFile()
	listUsername := regexp.MustCompile(helper.RegexCompileUsername()).FindAllString(snack, -1)

	defer file.Close()

	for _, username := range listUsername {
		includeStaging, _ := helper.IncludeArray(username, snackArray)

		if includeStaging == false {
			row := db.QueryRow("SELECT * FROM bandung_snack WHERE name='" + username + "'").Scan(&count)

			if row != sql.ErrNoRows {
				_, err := db.Exec("UPDATE bandung_snack SET status='belum' WHERE name='" + username + "'")
				helper.ErrorMessage(err)

				file.WriteString("\n- Kak " + username + " ora jadi bawa snack")
			} else {
				file.WriteString("\n- <code>" + username + "</code> siapa tuh? 👻")
			}

			snackArray = append(snackArray, username)
		}
	}
}
