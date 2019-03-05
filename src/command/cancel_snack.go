package command

import (
	"db"
	"helper"
	"message"
	"strings"
	"user"
)

func MatchCancelSnack() string {
	pattern := strings.HasPrefix(text_msg, helper.PrefixCommandCancelSnack())

	if pattern == true {
		GoToFunc = CancelSnack
	} else {
		return send_message
	}

	return GoToFunc()
}

func CancelSnack() string {
	var admin bool
	var snack string

	admin = user.IsAdmin(first_name)
	snack = helper.CheckEmptyUsername(text_msg)

	if admin {
		db.CancelSnack(text_msg)
	}

	if snack != "" {
		send_message = message.CancelSnack(snack, first_name)
	} else {
		send_message = message.EmptyUsername(first_name, base_command)
	}

	return send_message
}
