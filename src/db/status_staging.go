package db

import (
	"helper"
	"regexp"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func StatusStaging(staging string) {
	var staging_name, status, branch, username string

	db := DBConnection()
	file := helper.CreateFile()
	pattern := regexp.MustCompile(helper.RegexCompileStatusStaging())
	stg := pattern.FindAllString(staging, -1)

	defer file.Close()

	for _, match := range stg {
		stg_match := strings.Trim(match, " ")
		rows, err := db.Query("SELECT book_staging, book_status, book_branch, book_name FROM booking_staging where (book_staging='" + stg_match + "' or book_squad='" + stg_match + "')")
		helper.ErrorMessage(err)

		for rows.Next() {
			err = rows.Scan(&staging_name, &status, &branch, &username)
			helper.ErrorMessage(err)

			content := "<code>Staging" + staging_name + "</code> : <b>" + strings.ToUpper(status) + "</b>\n" + branch + "\n@" + username + "\n\n"

			file.WriteString(content)
		}

		err = rows.Err()
		helper.ErrorMessage(err)
	}
}
