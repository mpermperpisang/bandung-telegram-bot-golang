package command

import (
	"regexp"

	"github.com/bandung-telegram-bot-golang/src/database"
	"github.com/bandung-telegram-bot-golang/src/helper"
	"github.com/bandung-telegram-bot-golang/src/message"
)

func NewOnboarding(text, username string) string {
	var contentMessage string

	nickName, _ := regexp.MatchString(helper.RegexCompileNickName(), text)
	jobTitle, _ := regexp.MatchString(helper.RegexCompileJobTitle(), text)
	squad, _ := regexp.MatchString(helper.RegexCompileSquad(), text)
	lastJob, _ := regexp.MatchString(helper.RegexCompileLastJob(), text)
	status, _ := regexp.MatchString(helper.RegexCompileStatus(), text)
	hobby, _ := regexp.MatchString(helper.RegexCompileHobby(), text)
	motto, _ := regexp.MatchString(helper.RegexCompileMotto(), text)

	if !nickName || !jobTitle || !squad || !lastJob || !status || !hobby || !motto {
		contentMessage = message.InvalidOnboarding()
	} else {
		database.UpdateOnboarding(username)
		contentMessage = message.ValidOnboarding()
	}

	return contentMessage
}
