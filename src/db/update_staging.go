package db

import (
	"helper"
	"regexp"

	_ "github.com/go-sql-driver/mysql"
)

func UpdateStaging(staging string) {
	var stgArray []string

	db := DBConnection()
	file := helper.CreateFile()
	stgNumber := regexp.MustCompile(helper.RegexCompileStagingNumber()).FindAllString(staging, -1)
	stgSquad := regexp.MustCompile(helper.RegexCompileStagingSquad()).FindString(staging)

	defer file.Close()

	file.WriteString("<b>" + stgSquad + "</b>" + " :")

	for _, list := range stgNumber {
		includeStaging, _ := helper.IncludeArray(list, stgArray)

		if includeStaging == false {
			_, err := db.Exec("UPDATE booking_staging SET book_squad='" + stgSquad + "' WHERE book_staging='" + list + "'")
			helper.ErrorMessage(err)

			file.WriteString("\n- staging" + list + ".vm")

			stgArray = append(stgArray, list)
		}
	}
}
