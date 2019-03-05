package db

import (
	"database/sql"
	"helper"
	"regexp"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func HolidaySnack(snack string) {
	var snackArray []string
	var count int

	db := DBConnection()
	file := helper.CreateFile()
	listUsername := regexp.MustCompile(helper.RegexCompileUsername()).FindAllString(snack, -1)

	defer file.Close()

	for _, username := range listUsername {
		includeStaging, _ := helper.IncludeArray(username, snackArray)

		if includeStaging == false {
			if username == "@all" {
				_, err := db.Exec("UPDATE bandung_snack SET status='libur' WHERE day='" + strings.ToLower(helper.DayNow()) + "' and name<>''")
				helper.ErrorMessage(err)

				file.WriteString("\nSelamat hari libur berjamaah, Kak")
			} else {
				row := db.QueryRow("SELECT * FROM bandung_snack WHERE name='" + username + "'").Scan(&count)

				if row != sql.ErrNoRows {
					_, err := db.Exec("UPDATE bandung_snack SET status='libur' WHERE name='" + username + "'")
					helper.ErrorMessage(err)

					file.WriteString("\n- Kak <code>" + username + "</code> izin libur")
				} else {
					file.WriteString("\n- <code>" + username + "</code> siapa tuh? 👻")
				}
			}

			snackArray = append(snackArray, username)
		}
	}
}
