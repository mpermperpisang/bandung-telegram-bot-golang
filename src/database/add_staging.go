package database

import (
	"regexp"

	"github.com/bandung-telegram-bot-golang/src/helper"
	_ "github.com/go-sql-driver/mysql"
)

func AddStaging(staging string) {
	var stgArray []string

	db := DBConnection()
	file := helper.CreateFile()
	defer file.Close()
	stgNumber := regexp.MustCompile(helper.RegexCompileStagingNumber()).FindAllString(staging, -1)
	stgSquad := regexp.MustCompile(helper.RegexCompileStagingSquad()).FindString(staging)

	for _, list := range stgNumber {
		includeStaging, _ := helper.IncludeArray(list, stgArray)

		if !includeStaging {
			_, err := db.Exec("INSERT INTO booking_staging VALUES ('" + list + "', 'BookByUsername', '0', 'Branch', 'done', '" + stgSquad + "')")
			helper.ErrorMessage(err)

			if err != nil {
				file.WriteString("\n- <b>staging" + list + ".vm</b> sudah terdaftar")
			} else {
				file.WriteString("\n- <b>staging" + list + ".vm</b> berhasil ditambahkan")
			}

			stgArray = append(stgArray, list)
		}
	}
}
