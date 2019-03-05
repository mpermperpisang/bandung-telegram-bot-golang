package command

import (
	"db"
	"helper"
	"message"
	"strings"
	"user"
)

func MatchAddSnack() string {
	pattern := strings.HasPrefix(text_msg, helper.PrefixCommandAddSnack())

	if pattern == true {
		GoToFunc = AddSnack
	} else {
		return send_message
	}

	return GoToFunc()
}

func AddSnack() string {
	var admin bool
	var snack string

	admin = user.IsAdmin(user_name)
	snack = helper.CheckEmptyDayUsername(text_msg)

	if admin {
		db.AddSnack(snack)
	}

	if snack != "" {
		send_message = message.AddSnack(snack, user_name)
	} else {
		send_message = message.EmptyDayUsername(user_name, base_command)
	}

	return send_message
}
