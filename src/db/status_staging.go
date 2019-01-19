package db

import (
	"fmt"
	"helper"
	"regexp"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func StatusStaging(staging string) {
	db := DBConnection()
	file := helper.CreateFile()
	pattern := regexp.MustCompile(helper.RegexCompileStatusStaging())
	stg := pattern.FindAllString(staging, -1)

	defer file.Close()

	for _, match := range stg {
		stg_match := strings.Trim(match, " ")
		rows, err := db.Query("SELECT book_staging, book_status, book_branch, book_name FROM booking_staging where (book_staging='" + stg_match + "' or book_squad='" + stg_match + "')")

		if err != nil {
			panic(err)
		}

		for rows.Next() {
			var staging_name string
			var status string
			var branch string
			var user_name string

			err = rows.Scan(&staging_name, &status, &branch, &user_name)

			if err != nil {
				fmt.Println(err)
				panic(err)
			}

			content := "<code>Staging" + staging_name + "</code> : <b>" + strings.ToUpper(status) + "</b>\n" + branch + "\n@" + user_name + "\n\n"

			file.WriteString(content)
		}

		err = rows.Err()

		if err != nil {
			panic(err)
		}
	}
}
