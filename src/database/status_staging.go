package database

import (
	"regexp"
	"strings"

	"github.com/bandung-telegram-bot-golang/src/helper"
	_ "github.com/go-sql-driver/mysql"
)

func StatusStaging(staging string) {
	var stgNumber, status, branch, username, user string
	var stgArray, stgStatus []string

	db := DBConnection()
	file := helper.CreateFile()
	stg := regexp.MustCompile(helper.RegexCompileStatusStaging()).FindAllString(staging, -1)

	defer file.Close()

	for _, list := range stg {
		includeStaging, _ := helper.IncludeArray(list, stgArray)

		if !includeStaging {
			stgList := strings.Trim(list, " ")
			rows, err := db.Query("SELECT book_staging, book_status, book_branch, book_name FROM booking_staging WHERE (book_staging='" + stgList + "' or book_squad='" + stgList + "')")
			helper.ErrorMessage(err)

			for rows.Next() {
				err = rows.Scan(&stgNumber, &status, &branch, &username)
				helper.ErrorMessage(err)

				if status == "done" {
					user = "<code>@" + username + "</code>"
				} else {
					user = "@" + username
				}

				content := "<code>Staging" + stgNumber + "</code> : <b>" + strings.ToUpper(status) + "</b>\n" + branch + "\n" + user + "\n\n"
				includeNumber, _ := helper.IncludeArray(stgNumber, stgStatus)

				if !includeNumber {
					file.WriteString(content)

					stgStatus = append(stgStatus, stgNumber)
				}
			}

			stgArray = append(stgArray, list)

			err = rows.Err()
			helper.ErrorMessage(err)
		}
	}
}
