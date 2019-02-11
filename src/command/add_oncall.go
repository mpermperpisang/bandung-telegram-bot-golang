package command

import (
	"helper"
	"math/rand"
	"message"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func MatchAddOnCall() string {
	pattern := strings.HasPrefix(text_msg, helper.PrefixCommandAddOnCall())

	if pattern == true {
		admin := []string{os.Getenv("ADMIN1"), os.Getenv("ADMIN2"), os.Getenv("ADMIN3"), os.Getenv("ADMIN4")}
		includeAdmin, _ := helper.IncludeArray(first_name, admin)

		if includeAdmin == true {
			GoToFunc = AddOnCall
		}
	} else {
		return send_message
	}

	return GoToFunc()
}

func AddOnCall() string {
	var content, oncall string
	var column = -1
	var sheet = helper.GoogleSheet()

	t := time.Now()
	year, _ := strconv.Atoi(t.Format("2006"))

	oncall = helper.CheckEmptyBackEndOnCall(text_msg)

	if oncall == "" {
		send_message = message.EmptyOncallUsername(first_name)
	} else {

		if sheet == nil {
			send_message = message.EmptyTabSheet(strconv.Itoa(year))
		} else {
			pattern := regexp.MustCompile(helper.RegexCompileBackEndOnCall())
			backend := pattern.FindAllString(text_msg, -1)

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

			send_message = message.AddOnCall(t.Format("2006"))
		}
	}

	return send_message
}
