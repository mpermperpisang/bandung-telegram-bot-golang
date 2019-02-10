package command

import (
	"helper"
	"math/rand"
	"message"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func MatchAddOnCall() string {
	pattern := strings.HasPrefix(text_msg, helper.PrefixCommandAddOnCall())

	if pattern == true {
		GoToFunc = AddOnCall
	} else {
		return send_message
	}

	return GoToFunc()
}

func AddOnCall() string {
	var z string
	var k = -1
	var t = time.Now()
	var sheet = helper.GoogleSheet()

	pattern := regexp.MustCompile(helper.RegexCompileBackEndOnCall())
	backend := pattern.FindAllString(text_msg, -1)
	year, _ := strconv.Atoi(t.Format("2006"))

	for j := 1; j <= 12; j++ {
		lastday := time.Date(year, time.Month(j+1), 0, 0, 0, 0, 0, time.UTC)
		value, _ := strconv.Atoi(lastday.Format("02"))
		index := rand.Perm(len(backend))

		for x := 0; x < 10; x++ {
			for _, match := range index {

				for i := 0; i < 2; i++ {
					one := time.Date(year, time.Month(j), k+1, 23, 0, 0, 0, time.UTC)
					two := time.Date(year, time.Month(j), 1, 23, 0, 0, 0, time.UTC)

					if k+1 < value {
						if int(one.Weekday()) < 5 && int(one.Weekday()) > 0 {
							k = k + 1
						} else {
							if int(two.Weekday()) == 0 {
								if k == -1 {
									k = k + -1
								}
							}
							k = k + 3
						}

						if k >= value {
							z = ""
						} else {
							z = backend[match]
						}

						sheet.Update(k, j-1, z)
					}
				}
			}
		}
		k = -1
	}

	err := sheet.Synchronize()
	helper.ErrorMessage(err)

	send_message = message.AddOnCall(t.Format("2006"))
	return send_message
}
