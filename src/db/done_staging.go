package db

import (
	"database/sql"
	"helper"
	"regexp"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func DoneStaging(staging, first_name string, user_id int) {
	var count int

	db := DBConnection()
	file := helper.CreateFile()
	stgNumber := regexp.MustCompile(helper.RegexCompileStatusStaging()).FindString(staging)

	defer file.Close()

	stgList := strings.Trim(stgNumber, " ")
	row := db.QueryRow("SELECT book_status FROM booking_staging WHERE book_staging='" + stgList + "'").Scan(&count)
	_, err := db.Exec("UPDATE booking_staging SET book_status='done' WHERE book_staging='" + stgList + "'")
	helper.ErrorMessage(err)

	if err != nil || row == sql.ErrNoRows {
		file.WriteString("staging yang belum ditambahkan ke daftar\nContoh : /add_staging squad_name staging_number1 staging_number2 staging_number3")
	} else {
		file.WriteString("staging" + stgList + ".vm")
	}
}
