package command

import (
	"strings"

	"github.com/bandung-telegram-bot-golang/src/db"
	"github.com/bandung-telegram-bot-golang/src/helper"
	"github.com/bandung-telegram-bot-golang/src/message"
)

func MatchBookStaging() string {
	pattern := strings.HasPrefix(textMsg, helper.PrefixCommandBookStaging())

	if pattern == true {
		GoToFunc = BookStaging
	} else {
		return "not match"
	}

	return GoToFunc()
}

func BookStaging() string {
	var staging string

	staging = helper.CheckEmptyStaging(textMsg)

	db.BookStaging((strings.ToUpper(staging)), userName, userID)

	if staging != "" {
		contentMessage = message.BookStaging(userName, strings.Trim(staging, " "))
	} else {
		contentMessage = message.EmptyBookDoneStaging(userName, baseCommand)
	}

	if strings.Contains(contentMessage, "memesan") {
		sendTo = msg.Chat
	} else {
		sendTo = sendToPrivate
	}

	return "success"
}
