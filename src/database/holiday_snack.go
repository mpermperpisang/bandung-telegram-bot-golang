package database

import (
	"database/sql"
	"regexp"
	"strings"

	"github.com/bandung-telegram-bot-golang/src/helper"
	_ "github.com/go-sql-driver/mysql"
)

func HolidaySnack(snack string) {
	var snackArray []string
	var count int

	db := DBConnection()
	file := helper.OpenFile()
	listUsername := regexp.MustCompile(helper.RegexCompileUsername()).FindAllString(snack, -1)

	for _, username := range listUsername {
		includeSnack, _ := helper.IncludeArray(username, snackArray)
		listSchedule := db.QueryRow("SELECT * FROM bandung_snack WHERE name='" + username + "' and day='" + strings.ToLower(helper.DayNow()) + "'").Scan(&count)

		if !includeSnack {
			if username == "@all" {
				_, err := db.Exec("UPDATE bandung_snack SET status='libur' WHERE day='" + strings.ToLower(helper.DayNow()) + "' and name<>''")
				helper.ErrorMessage(err)

				file.WriteString("\nSelamat hari libur berjamaah, Kak")
			} else {
				if listSchedule != sql.ErrNoRows {
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
	file.Close()
}
