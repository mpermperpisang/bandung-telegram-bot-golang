package command

import (
	"db"
	"helper"
	"message"
	"strings"
	"user"
)

func MatchDeleteAdminSnack() string {
	pattern := strings.HasPrefix(text_msg, helper.PrefixCommandDeleteAdminSnack())

	if pattern == true {
		GoToFunc = DeleteAdminSnack
	} else {
		return send_message
	}

	return GoToFunc()
}

func DeleteAdminSnack() string {
	var admin bool
	var snack string

	admin = user.IsAdmin(user_name)
	snack = helper.CheckEmptyUsername(text_msg)

	if admin {
		db.DeleteAdminSnack(text_msg)
	}

	if snack != "" {
		send_message = message.DeleteAdminSnack(snack, user_name)
	} else {
		send_message = message.EmptyUsername(user_name, base_command)
	}

	return send_message
}
