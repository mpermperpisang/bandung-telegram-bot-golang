package database

import (
	"database/sql"
	"strconv"
	"strings"

	"github.com/bandung-telegram-bot-golang/src/helper"
	_ "github.com/go-sql-driver/mysql"
)

func DoneSnack(username string, id int) {
	var count int

	db := DBConnection()
	file := helper.CreateFile()

	defer file.Close()

	rowSchedule := db.QueryRow("SELECT * FROM bandung_snack WHERE name='@" + username + "' and day='" + strings.ToLower(helper.DayNow()) + "' and status='belum'").Scan(&count)

	if rowSchedule == sql.ErrNoRows {
		file.WriteString("\n- <code>" + username + "</code> siapa tuh? 👻")
	} else {
		_, err := db.Exec("UPDATE bandung_snack SET status='sudah', from_id='" + strconv.Itoa(id) + "' WHERE name='@" + username + "' and day='" + strings.ToLower(helper.DayNow()) + "' and status='belum'")
		helper.ErrorMessage(err)

		file.WriteString("\nSelamat menggendutkan diri, kawan-kawan\n😈")
	}
}
