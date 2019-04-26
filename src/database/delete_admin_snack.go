package database

import (
	"database/sql"
	"regexp"

	"github.com/bandung-telegram-bot-golang/src/helper"
	_ "github.com/go-sql-driver/mysql"
)

func DeleteAdminSnack(username string) {
	var snackArray []string
	var count int

	db := DBConnection()
	file := helper.OpenFile()
	listUsername := regexp.MustCompile(helper.RegexCompileUsername()).FindAllString(username, -1)

	for _, username := range listUsername {
		includeSnack, _ := helper.IncludeArray(username, snackArray)

		if !includeSnack {
			listAdmin := db.QueryRow("SELECT * FROM admin_snack WHERE adm_username='" + username + "'").Scan(&count)

			if listAdmin != sql.ErrNoRows {
				_, err := db.Exec("DELETE FROM admin_snack WHERE adm_username='" + username + "'")
				helper.ErrorMessage(err)

				file.WriteString("\n- " + username + " sudah dihapus dari daftar admin snack")
			} else {
				file.WriteString("\n- <code>" + username + "</code> siapa tuh? 👻")
			}

			snackArray = append(snackArray, username)
		}
	}
	file.Close()
}
