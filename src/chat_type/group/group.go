package group

import (
	"command"
	"helper"
	"strings"
)

func SendMessage(text, username string) string {
	var send_message string

	if strings.HasPrefix(text, "/status_staging") {
		send_message = command.StatusStaging(text, username)
	} else {
		send_message = helper.Command()
	}

	return send_message
}
