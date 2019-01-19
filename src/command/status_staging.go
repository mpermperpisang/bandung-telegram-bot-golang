package command

import (
	"db"
	"helper"
	"message"
	"regexp"
	"strings"
)

var Username string

func StatusStaging(text, username string) string {
	Username = username

	return CheckEmptyStaging(text)
}

func CheckEmptyStaging(text string) string {
	var msg_status string

	pattern := regexp.MustCompile(helper.RegexCompileCheckEmptyStaging())
	staging := pattern.FindString(text)

	db.StatusStaging(strings.ToUpper(staging))

	if staging != "" {
		msg_status = message.StagingStatus(staging)
	} else {
		msg_status = message.EmptyStaging(Username)
	}

	return msg_status
}
