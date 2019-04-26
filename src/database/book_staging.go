package database

import (
	"database/sql"
	"regexp"
	"strconv"
	"strings"

	"github.com/bandung-telegram-bot-golang/src/helper"
	"github.com/bandung-telegram-bot-golang/src/message"
	_ "github.com/go-sql-driver/mysql"
)

func BookStaging(staging, username string, user_id int) {
	var count int
	var userBook string

	db := DBConnection()
	file := helper.OpenFile()
	stgNumber := regexp.MustCompile(helper.RegexCompileStatusStaging()).FindString(staging)
	stgList := strings.Trim(stgNumber, " ")
	stgCount := db.QueryRow("SELECT * FROM booking_staging WHERE book_staging='" + stgList + "'").Scan(&count)
	bookCount := db.QueryRow("SELECT * FROM booking_staging WHERE book_staging='" + stgList + "' and book_name='" + username + "' and book_status='booked'").Scan(&count)
	doneCount := db.QueryRow("SELECT * FROM booking_staging WHERE book_staging='" + stgList + "' and book_status='done'").Scan(&count)
	doneExist, _ := db.Query("SELECT book_name FROM booking_staging WHERE book_staging='" + stgList + "'")

	if stgCount != sql.ErrNoRows {
		if bookCount == sql.ErrNoRows {
			if doneCount == sql.ErrNoRows {
				for doneExist.Next() {
					err := doneExist.Scan(&userBook)
					helper.ErrorMessage(err)

					file.WriteString("<b>staging" + stgList + ".vm</b> masih dibooking sama Kak @" + userBook + "")
				}
			} else {
				_, err := db.Exec("UPDATE booking_staging SET book_name='" + username + "', book_from_id=" + strconv.Itoa(user_id) + ", book_status='booked' WHERE book_staging='" + stgList + "'")
				helper.ErrorMessage(err)

				file.WriteString("<b>staging" + stgList + ".vm</b>\nTolong yang lain jangan ada yang otak-atik yaa")
			}
		} else {
			file.WriteString("Kamu masih booking <b>staging" + stgList + ".vm</b> kok, Kak 😊")
		}
	} else {
		file.WriteString(message.ExampleCommandStaging("/add_staging"))
	}

	file.Close()
}
