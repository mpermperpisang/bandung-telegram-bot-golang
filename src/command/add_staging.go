package command

import (
	"db"
	"helper"
	"message"
	"strings"
)

func MatchAddStaging() string {
	pattern := strings.HasPrefix(text_msg, helper.PrefixCommandAddStaging())

	if pattern == true {
		GoToFunc = AddStaging
	} else {
		return send_message
	}

	return GoToFunc()
}

func AddStaging() string {
	var staging string

	staging = helper.CheckEmptySquadStaging(text_msg)

	db.AddStaging(strings.ToUpper(staging))

	if staging != "" {
		send_message = message.AddStaging(strings.Trim(staging, " "))
	} else {
		send_message = message.EmptySquadStaging(first_name, base_command)
	}

	return send_message
}
