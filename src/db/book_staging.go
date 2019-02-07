package db

import (
	"database/sql"
	"helper"
	"regexp"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func BookStaging(staging, first_name string, user_id int) {
	var count int

	db := DBConnection()
	file := helper.CreateFile()
	stg := regexp.MustCompile(helper.RegexCompileStatusStaging()).FindString(staging)

	defer file.Close()

	stgMatch := strings.Trim(stg, " ")
	row := db.QueryRow("SELECT book_status FROM booking_staging WHERE book_staging='" + stgMatch + "'").Scan(&count)
	_, err := db.Exec("UPDATE booking_staging SET book_name='" + first_name + "', book_from_id=" + strconv.Itoa(user_id) + ", book_status='booked' WHERE book_staging='" + stgMatch + "'")
	helper.ErrorMessage(err)

	if err != nil || row == sql.ErrNoRows {
		file.WriteString("staging yang belum ditambahkan ke daftar\nContoh : /add_staging squad_name staging_number1 staging_number2 staging_number3")
	} else {
		file.WriteString("staging" + stgMatch + ".vm\nTolong yang lain jangan ada yang otak-atik yaa")
	}
}
