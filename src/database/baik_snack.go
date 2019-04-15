package database

import (
	"github.com/bandung-telegram-bot-golang/src/helper"
	_ "github.com/go-sql-driver/mysql"
)

func BaikSnack(username string) {
	var name, count string

	db := DBConnection()
	file := helper.CreateFile()
	defer file.Close()
	listBaik, _ := db.Query("SELECT name, count FROM bandung_snack WHERE count<>0 ORDER BY count DESC LIMIT 25;")

	for listBaik.Next() {
		err := listBaik.Scan(&name, &count)
		helper.ErrorMessage(err)

		file.WriteString("- " + name + " bawa snack sebanyak " + count + " kali\n")
	}
}
