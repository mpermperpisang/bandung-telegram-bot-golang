package database

import (
	"database/sql"
	"strconv"
	"strings"

	"github.com/bandung-telegram-bot-golang/src/helper"
	_ "github.com/go-sql-driver/mysql"
)

func DoneSnack(username string, id int) {
	var count, countDone int

	db := DBConnection()
	file := helper.CreateFile()
	defer file.Close()
	snackCount, _ := db.Query("SELECT count FROM bandung_snack WHERE name='@" + username + "' and day='" + strings.ToLower(helper.DayNow()) + "' and status='belum'")
	rowSchedule := snackCount.Scan(&count)

	if rowSchedule == sql.ErrNoRows {
		file.WriteString("\n- <code>" + username + "</code> siapa tuh? 👻")
	} else {
		for snackCount.Next() {
			err := snackCount.Scan(&countDone)
			helper.ErrorMessage(err)

			countDone++
			_, err = db.Exec("UPDATE bandung_snack SET status='sudah', from_id='" + strconv.Itoa(id) + "', count=" + strconv.Itoa(countDone) + " WHERE name='@" + username + "' and day='" + strings.ToLower(helper.DayNow()) + "' and status='belum'")
			helper.ErrorMessage(err)

			file.WriteString("\nSelamat menggendutkan diri, kawan-kawan\n😈")
		}
	}
}
