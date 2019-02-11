package command

import (
	"db"
	"helper"
	"message"
	"strings"
)

func MatchStatusStaging() string {
	pattern := strings.HasPrefix(text_msg, helper.PrefixCommandStatusStaging())

	if pattern == true {
		GoToFunc = StatusStaging
	} else {
		return send_message
	}

	return GoToFunc()
}

func StatusStaging() string {
	var staging string

	staging = helper.CheckEmptyStaging(text_msg)

	db.StatusStaging(strings.ToUpper(staging))

	if staging != "" {
		send_message = message.StagingStatus(strings.Trim(staging, " "))
	} else {
		send_message = message.EmptyStaging(first_name)
	}

	return send_message
}
