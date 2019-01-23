package db

import (
	"helper"
	"regexp"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func BookStaging(staging, first_name string, user_id int) {
	db := DBConnection()
	file := helper.CreateFile()
	stg := regexp.MustCompile(helper.RegexCompileStatusStaging()).FindString(staging)

	defer file.Close()

	stg_match := strings.Trim(stg, " ")
	_, err := db.Exec("UPDATE booking_staging set book_name='" + first_name + "', book_from_id=" + strconv.Itoa(user_id) + ", book_status='booked' where book_staging='" + stg_match + "'")
	helper.ErrorMessage(err)

	file.WriteString("staging" + stg_match + ".vm")
}
