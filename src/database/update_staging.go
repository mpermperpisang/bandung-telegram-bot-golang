package database

import (
	"database/sql"
	"regexp"

	"github.com/bandung-telegram-bot-golang/src/helper"
	_ "github.com/go-sql-driver/mysql"
)

func UpdateStaging(staging string) {
	var stgArray []string
	var count int

	db := DBConnection()
	file := helper.CreateFile()
	stgNumber := regexp.MustCompile(helper.RegexCompileStagingNumber()).FindAllString(staging, -1)
	stgSquad := regexp.MustCompile(helper.RegexCompileStagingSquad()).FindString(staging)

	file.WriteString("<b>" + stgSquad + "</b>" + " :")

	for _, list := range stgNumber {
		includeStaging, _ := helper.IncludeArray(list, stgArray)

		if !includeStaging {
			rowStaging := db.QueryRow("SELECT * FROM booking_staging WHERE book_staging='" + list + "'").Scan(&count)

			if rowStaging == sql.ErrNoRows {
				file.WriteString("\n- staging" + list + ".vm tidak ditemukan")
			} else {
				_, err := db.Exec("UPDATE booking_staging SET book_squad='" + stgSquad + "' WHERE book_staging='" + list + "'")
				helper.ErrorMessage(err)

				file.WriteString("\n- staging" + list + ".vm berhasil diubah")
			}

			stgArray = append(stgArray, list)
		}
		file.Close()
	}
}
