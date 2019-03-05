package command

import (
	"db"
	"helper"
	"message"
	"strings"
	"user"
)

func MatchHolidaySnack() string {
	pattern := strings.HasPrefix(text_msg, helper.PrefixCommandHolidaySnack())

	if pattern == true {
		GoToFunc = HolidaySnack
	} else {
		return send_message
	}

	return GoToFunc()
}

func HolidaySnack() string {
	var admin bool
	var snack string

	admin = user.IsAdmin(user_name)
	snack = helper.CheckEmptyUsername(text_msg)

	if admin {
		db.HolidaySnack(text_msg)
	}

	if snack != "" {
		send_message = message.HolidaySnack(snack, user_name)
	} else {
		send_message = message.EmptyUsername(user_name, base_command)
	}

	return send_message
}
