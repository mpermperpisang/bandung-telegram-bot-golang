package database

import (
	"database/sql"
	"regexp"
	"strings"

	"github.com/bandung-telegram-bot-golang/src/helper"
	"github.com/bandung-telegram-bot-golang/src/message"
	_ "github.com/go-sql-driver/mysql"
)

func DoneStaging(staging, first_name string, user_id int) {
	var count int

	db := DBConnection()
	file := helper.OpenFile()
	stgNumber := regexp.MustCompile(helper.RegexCompileStatusStaging()).FindString(staging)
	stgList := strings.Trim(stgNumber, " ")
	row := db.QueryRow("SELECT book_status FROM booking_staging WHERE book_staging='" + stgList + "'").Scan(&count)
	_, err := db.Exec("UPDATE booking_staging SET book_status='done' WHERE book_staging='" + stgList + "'")
	helper.ErrorMessage(err)

	if err != nil || row == sql.ErrNoRows {
		file.WriteString(message.ExampleCommandStaging("/add_staging"))
	} else {
		file.WriteString("<b>staging" + stgList + ".vm</b>")
	}

	file.Close()
}
