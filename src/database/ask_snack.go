package database

import (
	"strings"

	"github.com/bandung-telegram-bot-golang/src/helper"
	_ "github.com/go-sql-driver/mysql"
)

func AskSnack(username string) {
	var userSnackID string

	db := DBConnection()
	file := helper.CreateFile()
	defer file.Close()
	listDone, err := db.Query("SELECT from_id FROM bandung_snack WHERE day='" + strings.ToLower(helper.DayNow()) + "' and status='sudah' and name<>'@" + username + "'")

	for listDone.Next() {
		err = listDone.Scan(&userSnackID)
		helper.ErrorMessage(err)

		file.WriteString(userSnackID + "\n")
	}
}
