package command

import (
	"helper"
	"message"
	"strconv"
	"strings"
	"time"
)

func MatchOnCall() string {
	pattern := strings.HasPrefix(text_msg, helper.PrefixCommandOnCall())

	if pattern == true {
		GoToFunc = OnCall
	} else {
		return send_message
	}

	return GoToFunc()
}

func OnCall() string {
	var imonth, idate int
	var sheet = helper.GoogleSheet()

	t := time.Now()
	month, _ := strconv.Atoi(t.Format("01"))
	date, _ := strconv.Atoi(t.Format("02"))
	year, _ := strconv.Atoi(t.Format("2006"))

	if sheet == nil || sheet.Columns[imonth][idate].Value == "" {
		send_message = message.EmptyOnCall(strconv.Itoa(year))
	} else {
		imonth = int(month) - 1
		idate = int(date) - 1
		contentsheet := sheet.Columns[imonth][idate].Value

		if contentsheet == "" {
			send_message = message.HolidayOnCall()
		} else {
			send_message = message.OnCall(sheet.Columns[imonth][idate].Value)
		}
	}

	return send_message
}
