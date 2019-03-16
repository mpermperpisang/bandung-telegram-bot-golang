package db

import (
	"database/sql"
	"regexp"
	"strconv"
	"strings"

	"github.com/bandung-telegram-bot-golang/src/helper"
	_ "github.com/go-sql-driver/mysql"
)

func BookStaging(staging, first_name string, user_id int) {
	var count int

	db := DBConnection()
	file := helper.CreateFile()
	stgNumber := regexp.MustCompile(helper.RegexCompileStatusStaging()).FindString(staging)

	defer file.Close()

	stgList := strings.Trim(stgNumber, " ")
	row := db.QueryRow("SELECT book_status FROM booking_staging WHERE book_staging='" + stgList + "'").Scan(&count)
	_, err := db.Exec("UPDATE booking_staging SET book_name='" + first_name + "', book_from_id=" + strconv.Itoa(user_id) + ", book_status='booked' WHERE book_staging='" + stgList + "'")
	helper.ErrorMessage(err)

	if err != nil || row == sql.ErrNoRows {
		file.WriteString("staging yang belum ditambahkan ke daftar\nContoh : <code>/add_staging squad_name staging_number1 staging_number2 staging_number3</code> dan seterusnya")
	} else {
		file.WriteString("staging" + stgList + ".vm\nTolong yang lain jangan ada yang otak-atik yaa")
	}
}
