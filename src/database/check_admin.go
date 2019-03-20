package database

import (
	"database/sql"

	"github.com/bandung-telegram-bot-golang/src/helper"
	_ "github.com/go-sql-driver/mysql"
)

func CheckAdmin(username string) bool {
	var count int

	db := DBConnection()
	file := helper.CreateFile()

	defer file.Close()

	row := db.QueryRow("SELECT * FROM admin_snack WHERE adm_username='@" + username + "'").Scan(&count)

	if row == sql.ErrNoRows {
		return false
	} else {
		return true
	}
}
