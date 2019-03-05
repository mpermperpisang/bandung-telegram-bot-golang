package command

import (
	"db"
	"helper"
	"message"
	"strings"
)

func MatchListAdminSnack() string {
	pattern := strings.HasPrefix(text_msg, helper.PrefixCommandListAdminSnack())

	if pattern == true {
		GoToFunc = ListAdminSnack
	} else {
		return send_message
	}

	return GoToFunc()
}

func ListAdminSnack() string {
	db.ListAdminSnack()

	send_message = message.ListAdminSnack(user_name)

	return send_message
}
