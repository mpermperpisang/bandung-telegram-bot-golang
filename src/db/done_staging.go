package db

import (
	"helper"
	"regexp"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func DoneStaging(staging, first_name string, user_id int) {
	db := DBConnection()
	file := helper.CreateFile()
	stg := regexp.MustCompile(helper.RegexCompileStatusStaging()).FindString(staging)

	defer file.Close()

	stg_match := strings.Trim(stg, " ")
	_, err := db.Exec("UPDATE booking_staging set book_status='done' where book_staging='" + stg_match + "'")
	helper.ErrorMessage(err)

	file.WriteString("staging" + stg_match + ".vm")
}
