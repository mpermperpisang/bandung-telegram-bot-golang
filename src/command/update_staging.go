package command

import (
	"db"
	"helper"
	"message"
	"strings"
)

func MatchUpdateStaging() string {
	pattern := strings.HasPrefix(text_msg, helper.PrefixCommandUpdateStaging())

	if pattern == true {
		GoToFunc = UpdateStaging
	} else {
		return send_message
	}

	return GoToFunc()
}

func UpdateStaging() string {
	var staging string

	staging = helper.CheckEmptySquadStaging(text_msg)

	db.UpdateStaging(strings.ToUpper(staging))

	if staging != "" {
		send_message = message.UpdateStaging(strings.Trim(staging, " "))
	} else {
		send_message = message.EmptySquadStaging(first_name, base_command)
	}

	return send_message
}
