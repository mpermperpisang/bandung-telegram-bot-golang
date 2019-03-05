package db

import (
	"database/sql"
	"helper"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func DoneSnack(username string, id int) {
	var count int

	db := DBConnection()
	file := helper.CreateFile()

	defer file.Close()

	rowUsername := db.QueryRow("SELECT * FROM bandung_snack WHERE name='@" + username + "'").Scan(&count)
	rowSchedule := db.QueryRow("SELECT * FROM bandung_snack WHERE name='@" + username + "', day='" + strings.ToLower(helper.DayNow()) + "' and status='belum'").Scan(&count)

	if rowUsername != sql.ErrNoRows && rowSchedule == sql.ErrNoRows {
		_, err := db.Exec("UPDATE bandung_snack SET status='sudah', from_id='" + strconv.Itoa(id) + "' WHERE name='@" + username + "' and day='" + strings.ToLower(helper.DayNow()) + "' and status='belum'")
		helper.ErrorMessage(err)

		file.WriteString("\nSelamat menggendutkan diri, kawan-kawan\n😈")
	} else {
		file.WriteString("\n- <code>" + username + "</code> siapa tuh? 👻")
	}
}
