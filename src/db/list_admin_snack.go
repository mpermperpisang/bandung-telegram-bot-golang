package db

import (
	"github.com/bandung-telegram-bot-golang/src/helper"
	_ "github.com/go-sql-driver/mysql"
)

func ListAdminSnack() {
	var adminUsername string

	db := DBConnection()
	file := helper.CreateFile()

	defer file.Close()
	rows, err := db.Query("SELECT adm_username FROM admin_snack")

	for rows.Next() {
		err = rows.Scan(&adminUsername)
		helper.ErrorMessage(err)

		file.WriteString("- <code>" + adminUsername + "</code>\n")
	}
}
