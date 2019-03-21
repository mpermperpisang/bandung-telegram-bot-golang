package command

import (
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/bandung-telegram-bot-golang/src/helper"
	"github.com/bandung-telegram-bot-golang/src/message"
)

func MatchAddOnCall() string {
	pattern := strings.HasPrefix(textMsg, helper.PrefixCommandAddOnCall())
	admin := []string{os.Getenv("ADMIN1"), os.Getenv("ADMIN2"), os.Getenv("ADMIN3"), os.Getenv("ADMIN4")}
	includeAdmin, _ := helper.IncludeArray("@"+userName, admin)

	if pattern {
		if includeAdmin {
			GoToFunc = AddOnCall
		} else {
			contentMessage = message.OncallAdmin(userName, baseCommand, admin)
			sendTo = sendToPrivate
			return "not admin"
		}
	} else {
		return "not match"
	}

	return GoToFunc()
}

func AddOnCall() string {
	var content, oncall string

	var column = -1
	var sheet = helper.GoogleSheet(os.Getenv("SPREADSHEET_ONCALL"))

	t := time.Now()
	year, _ := strconv.Atoi(t.Format("2006"))
	oncall = helper.CheckEmptyUsername(textMsg)

	if oncall == "" {
		contentMessage = message.EmptyUsername(userName, baseCommand)
	} else {
		if sheet == nil {
			contentMessage = message.EmptyTabSheet(strconv.Itoa(year))
		} else {
			pattern := regexp.MustCompile(helper.RegexCompileUsername())
			backend := pattern.FindAllString(textMsg, -1)

			for row := 1; row <= 12; row++ {
				lastday := time.Date(year, time.Month(row+1), 0, 0, 0, 0, 0, time.UTC)
				enddate, _ := strconv.Atoi(lastday.Format("02"))
				index := rand.Perm(len(backend))

				for x := 0; x < 5; x++ {
					for _, match := range index {
						for i := 0; i < 2; i++ {
							one := time.Date(year, time.Month(row), column+1, 23, 0, 0, 0, time.UTC)
							two := time.Date(year, time.Month(row), 1, 23, 0, 0, 0, time.UTC)

							if column+1 < enddate {
								if int(one.Weekday()) < 5 && int(one.Weekday()) > 0 {
									column = column + 1
								} else {
									if int(two.Weekday()) == 0 {
										if column == -1 {
											column = column + -1
										}
									}
									column = column + 3
								}

								if column < enddate {
									content = backend[match]

									sheet.Update(column, row-1, content)
								}
							}
						}
					}
				}
				column = -1
			}

			err := sheet.Synchronize()
			helper.ErrorMessage(err)

			contentMessage = message.AddOnCall(t.Format("2006"))
		}
	}

	if strings.Contains(contentMessage, "Berhasil") {
		sendTo = msg.Chat
	} else {
		sendTo = sendToPrivate
	}

	return "success"
}
