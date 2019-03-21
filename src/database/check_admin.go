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
	listAdmin := db.QueryRow("SELECT * FROM admin_snack WHERE adm_username='@" + username + "'").Scan(&count)

	defer file.Close()

	if listAdmin == sql.ErrNoRows {
		return false
	} else {
		return true
	}
}
