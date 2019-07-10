package command

import (
	"strings"

	"github.com/bandung-telegram-bot-golang/src/database"
	"github.com/bandung-telegram-bot-golang/src/helper"
	"github.com/bandung-telegram-bot-golang/src/message"
	"gopkg.in/tucnak/telebot.v2"
)

func MatchDoneSnack() string {
	pattern := strings.HasPrefix(textMsg, helper.PrefixCommandDoneSnack())

	if pattern {
		GoToFunc = DoneSnack
	} else {
		return "not match"
	}

	return GoToFunc()
}

func DoneSnack() string {
	database.DoneSnack(userName, userID)

	contentMessage = message.DoneSnack(userName, firstName)

	if strings.Contains(contentMessage, "Yeay") {
		sendTo = sendToGroup
		bot.Send(sendToPrivate, message.PostToChannel(), telebot.ModeHTML)
	} else {
		sendTo = sendToPrivate
	}

	return "success"
}
