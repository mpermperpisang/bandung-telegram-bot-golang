package db

import (
	"helper"
	"regexp"

	_ "github.com/go-sql-driver/mysql"
)

func DeleteAdminSnack(username string) {
	var snackArray []string

	db := DBConnection()
	file := helper.CreateFile()
	listUsername := regexp.MustCompile(helper.RegexCompileUsername()).FindAllString(username, -1)

	defer file.Close()

	for _, username := range listUsername {
		includeStaging, _ := helper.IncludeArray(username, snackArray)

		if includeStaging == false {
			_, err := db.Exec("DELETE FROM admin_snack WHERE adm_username='" + username + "'")
			helper.ErrorMessage(err)

			if err != nil {
				file.WriteString("\n- <code>" + username + "</code> siapa tuh? 👻")
			} else {
				file.WriteString("\n- " + username + " sudah dihapus dari daftar admin snack")
			}

			snackArray = append(snackArray, username)
		}
	}
}
