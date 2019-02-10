package db

import (
	"helper"
	"regexp"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func StatusStaging(staging string) {
	var stagingName, status, branch, username, user string
	var stgArray []string

	db := DBConnection()
	file := helper.CreateFile()
	pattern := regexp.MustCompile(helper.RegexCompileStatusStaging())
	stg := pattern.FindAllString(staging, -1)

	defer file.Close()

	for _, match := range stg {
		includeStaging, _ := helper.IncludeArray(match, stgArray)
		if includeStaging == false {
			stgMatch := strings.Trim(match, " ")
			rows, err := db.Query("SELECT book_staging, book_status, book_branch, book_name FROM booking_staging WHERE (book_staging='" + stgMatch + "' or book_squad='" + stgMatch + "')")
			helper.ErrorMessage(err)

			for rows.Next() {
				err = rows.Scan(&stagingName, &status, &branch, &username)
				helper.ErrorMessage(err)

				if status == "done" {
					user = "<code>@" + username + "</code>"
				} else {
					user = "@" + username
				}

				content := "<code>Staging" + stagingName + "</code> : <b>" + strings.ToUpper(status) + "</b>\n" + branch + "\n" + user + "\n\n"

				file.WriteString(content)
			}

			stgArray = append(stgArray, match)

			err = rows.Err()
			helper.ErrorMessage(err)
		}
	}
}
