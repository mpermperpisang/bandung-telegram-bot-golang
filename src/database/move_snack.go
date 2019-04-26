package database

import (
	"database/sql"
	"regexp"
	"strings"

	"github.com/bandung-telegram-bot-golang/src/helper"
	_ "github.com/go-sql-driver/mysql"
)

func MoveSnack(snack string) {
	var snackArray []string
	var count int

	db := DBConnection()
	file := helper.CreateFile()
	listSchedule := regexp.MustCompile(helper.RegexCompileSnackSchedule()).FindAllString(snack, -1)

	for _, list := range listSchedule {
		snackDay := regexp.MustCompile(helper.RegexCompileSnackDay()).FindString(list)
		snackUsername := regexp.MustCompile(helper.RegexCompileUsername()).FindString(list)
		includeSnack, _ := helper.IncludeArray(snackUsername, snackArray)
		row := db.QueryRow("SELECT * FROM bandung_snack WHERE name='" + snackUsername + "'").Scan(&count)

		if !includeSnack {
			if row != sql.ErrNoRows {
				_, err := db.Exec("UPDATE bandung_snack SET day='" + strings.Trim(snackDay, " ") + "' WHERE name='" + snackUsername + "'")
				helper.ErrorMessage(err)

				file.WriteString("\n- " + snackUsername + " pindah sementara ke hari <b>" + helper.DayName(strings.Trim(snackDay, " ")) + "</b>")
			} else {
				file.WriteString("\n- <code>" + snackUsername + "</code> siapa tuh? 👻")
			}

			snackArray = append(snackArray, snackUsername)
		}
		file.Close()
	}
}
