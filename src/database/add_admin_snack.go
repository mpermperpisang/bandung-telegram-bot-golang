package database

import (
	"regexp"

	"github.com/bandung-telegram-bot-golang/src/helper"
)

func AddAdminSnack(username string) {
	var snackArray []string

	db := DBConnection()
	file := helper.CreateFile()
	listUsername := regexp.MustCompile(helper.RegexCompileUsername()).FindAllString(username, -1)

	defer file.Close()

	for _, username := range listUsername {
		includeSnack, _ := helper.IncludeArray(username, snackArray)

		if !includeSnack {
			_, err := db.Exec("INSERT INTO admin_snack (adm_username) VALUES ('" + username + "')")
			helper.ErrorMessage(err)

			if err != nil {
				file.WriteString("\n- <code>" + username + "</code> sudah jadi admin snack")
			} else {
				file.WriteString("\n- " + username + " sudah resmi jadi admin snack")
			}

			snackArray = append(snackArray, username)
		}
	}
}
