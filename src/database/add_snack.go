package database

import (
	"regexp"
	"strings"

	"github.com/bandung-telegram-bot-golang/src/helper"
	_ "github.com/go-sql-driver/mysql"
)

func AddSnack(snack string) {
	var snackArray []string

	db := DBConnection()
	file := helper.CreateFile()
	listSchedule := regexp.MustCompile(helper.RegexCompileSnackSchedule()).FindAllString(snack, -1)

	defer file.Close()

	for _, list := range listSchedule {
		snackDay := regexp.MustCompile(helper.RegexCompileSnackDay()).FindString(list)
		snackUsername := regexp.MustCompile(helper.RegexCompileUsername()).FindString(list)
		includeSnack, _ := helper.IncludeArray(snackUsername, snackArray)

		if !includeSnack {
			_, err := db.Exec("INSERT INTO bandung_snack VALUES ('" + strings.Trim(snackDay, " ") + "', '" + strings.Trim(snackDay, " ") + "', '" + snackUsername + "', 'sudah', '-')")
			helper.ErrorMessage(err)

			if err != nil {
				file.WriteString("\n- <code>" + snackUsername + "</code> sudah ada jadwal snack")
			} else {
				file.WriteString("\n- " + snackUsername + " bawa snack di hari <b>" + helper.DayName(strings.Trim(snackDay, " ")) + "</b>")
			}

			snackArray = append(snackArray, snackUsername)
		}
	}
}
