package database

import (
	"github.com/bandung-telegram-bot-golang/src/helper"
	_ "github.com/go-sql-driver/mysql"
)

func ListAdminSnack() {
	var adminUsername string

	db := DBConnection()
	file := helper.CreateFile()
	listAdmin, err := db.Query("SELECT adm_username FROM admin_snack")

	for listAdmin.Next() {
		err = listAdmin.Scan(&adminUsername)
		helper.ErrorMessage(err)

		file.WriteString("- <code>" + adminUsername + "</code>\n")
	}

	defer file.Close()
}
