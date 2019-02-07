package db

import (
	"helper"
	"regexp"

	_ "github.com/go-sql-driver/mysql"
)

func UpdateStaging(staging string) {
	db := DBConnection()
	file := helper.CreateFile()
	pattern := regexp.MustCompile(helper.RegexCompileStagingNumber())
	stgNumber := pattern.FindAllString(staging, -1)
	stgSquad := regexp.MustCompile(helper.RegexCompileStagingSquad()).FindString(staging)

	defer file.Close()

	for _, match := range stgNumber {
		_, err := db.Exec("UPDATE booking_staging SET book_squad='" + stgSquad + "' where book_staging='" + match + "'")
		helper.ErrorMessage(err)

		file.WriteString("\n- staging" + match + ".vm")
	}
}
