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
	pattern := regexp.MustCompile(helper.RegexCompileStagingNumber())
	stgNumber := pattern.FindAllString(staging, -1)
	stgSquad := regexp.MustCompile(helper.RegexCompileStagingSquad()).FindString(staging)

	defer file.Close()

	for _, match := range stgNumber {
		includeStaging, _ := helper.IncludeArray(match, stgArray)
		if includeStaging == false {
			_, err := db.Exec("INSERT INTO booking_staging VALUES ('" + match + "', 'book', '0', 'book', 'done', '" + stgSquad + "')")
			helper.ErrorMessage(err)

			if err != nil {
				file.WriteString("\n- <b>staging" + match + ".vm</b> sudah terdaftar")
			} else {
				file.WriteString("\n- <b>staging" + match + ".vm</b> berhasil ditambahkan")
			}

			stgArray = append(stgArray, match)
		}
	}
}
