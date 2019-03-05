package command

import (
	"db"
	"helper"
	"message"
	"strings"
	"user"
)

func MatchMoveSnack() string {
	pattern := strings.HasPrefix(text_msg, helper.PrefixCommandMoveSnack())

	if pattern == true {
		GoToFunc = MoveSnack
	} else {
		return send_message
	}

	return GoToFunc()
}

func MoveSnack() string {
	var admin bool
	var snack string

	admin = user.IsAdmin(first_name)
	snack = helper.CheckEmptyDayUsername(text_msg)

	if admin {
		db.MoveSnack(snack)
	}

	if snack != "" {
		send_message = message.MoveSnack(snack, first_name)
	} else {
		send_message = message.EmptyDayUsername(first_name, base_command)
	}

	return send_message
}
