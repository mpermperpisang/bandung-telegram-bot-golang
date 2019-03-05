package command

import (
	"db"
	"helper"
	"message"
	"strings"
	"user"
)

func MatchAddAdminSnack() string {
	pattern := strings.HasPrefix(text_msg, helper.PrefixCommandAddAdminSnack())

	if pattern == true {
		GoToFunc = AddAdminSnack
	} else {
		return send_message
	}

	return GoToFunc()
}

func AddAdminSnack() string {
	var admin bool
	var snack string

	admin = user.IsAdmin(user_name)
	snack = helper.CheckEmptyUsername(text_msg)

	if admin {
		db.AddAdminSnack(text_msg)
	}

	if snack != "" {
		send_message = message.AddAdminSnack(snack, user_name)
	} else {
		send_message = message.EmptyUsername(user_name, base_command)
	}

	return send_message
}
