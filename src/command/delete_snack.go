package command

import (
	"db"
	"helper"
	"message"
	"strings"
	"user"
)

func MatchDeleteSnack() string {
	pattern := strings.HasPrefix(text_msg, helper.PrefixCommandDeleteSnack())

	if pattern == true {
		GoToFunc = DeleteSnack
	} else {
		return send_message
	}

	return GoToFunc()
}

func DeleteSnack() string {
	var admin bool
	var snack string

	admin = user.IsAdmin(first_name)
	snack = helper.CheckEmptyUsername(text_msg)

	if admin {
		db.DeleteSnack(text_msg)
	}

	if snack != "" {
		send_message = message.DeleteSnack(snack, first_name)
	} else {
		send_message = message.EmptyUsername(first_name, base_command)
	}

	return send_message
}
