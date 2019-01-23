package db

import (
	"helper"
	"regexp"

	_ "github.com/go-sql-driver/mysql"
)

func AddStaging(staging string) {
	db := DBConnection()
	file := helper.CreateFile()
	pattern := regexp.MustCompile(helper.RegexCompileStagingNumber())
	stg_number := pattern.FindAllString(staging, -1)
	stg_squad := regexp.MustCompile(helper.RegexCompileStagingSquad()).FindString(staging)

	defer file.Close()

	for _, match := range stg_number {
		_, err := db.Exec("INSERT INTO booking_staging VALUES ('" + match + "', 'book', '0', 'book', 'done', '" + stg_squad + "')")
		helper.ErrorMessage(err)

		file.WriteString("\n- staging" + match + ".vm")
	}
}
