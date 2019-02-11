package db

import (
	"helper"
	"regexp"

	_ "github.com/go-sql-driver/mysql"
)

func AddStaging(staging string) {
	var stgArray []string

	db := DBConnection()
	file := helper.CreateFile()
	stgNumber := regexp.MustCompile(helper.RegexCompileStagingNumber()).FindAllString(staging, -1)
	stgSquad := regexp.MustCompile(helper.RegexCompileStagingSquad()).FindString(staging)

	defer file.Close()

	for _, list := range stgNumber {
		includeStaging, _ := helper.IncludeArray(list, stgArray)

		if includeStaging == false {
			_, err := db.Exec("INSERT INTO booking_staging VALUES ('" + list + "', 'book', '0', 'book', 'done', '" + stgSquad + "')")
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
