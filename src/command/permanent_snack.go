package command

import (
	"db"
	"helper"
	"message"
	"strings"
	"user"
)

func MatchPermanentSnack() string {
	pattern := strings.HasPrefix(text_msg, helper.PrefixCommandPermanentSnack())

	if pattern == true {
		GoToFunc = PermanentSnack
	} else {
		return send_message
	}

	return GoToFunc()
}

func PermanentSnack() string {
	var admin bool
	var snack string

	admin = user.IsAdmin(first_name)
	snack = helper.CheckEmptyDayUsername(text_msg)

	if admin {
		db.PermanentSnack(snack)
	}

	if snack != "" {
		send_message = message.PermanentSnack(snack, first_name)
	} else {
		send_message = message.EmptyDayUsername(first_name, base_command)
	}

	return send_message
}
