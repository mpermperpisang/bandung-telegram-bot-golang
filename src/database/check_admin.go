package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func CheckAdmin(username string) bool {
	var count int

	db := DBConnection()
	listAdmin := db.QueryRow("SELECT * FROM admin_snack WHERE adm_username='@" + username + "'").Scan(&count)

	if listAdmin == sql.ErrNoRows {
		return false
	} else {
		return true
	}
}
