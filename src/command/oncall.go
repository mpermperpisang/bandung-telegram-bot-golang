package command

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/bandung-telegram-bot-golang/src/helper"
	"github.com/bandung-telegram-bot-golang/src/message"
)

func MatchOnCall() string {
	pattern := strings.HasPrefix(textMsg, helper.PrefixCommandOnCall())

	if pattern {
		GoToFunc = OnCall
	} else {
		return "not match"
	}

	return GoToFunc()
}

func OnCall() string {
	var imonth, idate int
	var sheet = helper.GoogleSheet(os.Getenv("SPREADSHEET_ONCALL"))

	t := time.Now()
	month, _ := strconv.Atoi(t.Format("01"))
	date, _ := strconv.Atoi(t.Format("02"))
	year, _ := strconv.Atoi(t.Format("2006"))

	if sheet == nil || sheet.Columns[imonth][idate].Value == "" {
		contentMessage = message.EmptyOnCall(strconv.Itoa(year))
	} else {
		imonth = int(month) - 1
		idate = int(date) - 1
		contentsheet := sheet.Columns[imonth][idate].Value

		if contentsheet == "" {
			contentMessage = message.HolidayOnCall()
		} else {
			contentMessage = message.OnCall(sheet.Columns[imonth][idate].Value)
		}
	}

	if strings.Contains(contentMessage, "Semangat") {
		sendTo = msg.Chat
	} else {
		sendTo = sendToPrivate
	}

	return "success"
}
