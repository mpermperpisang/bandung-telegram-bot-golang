package command

import (
	"gopkg.in/tucnak/telebot.v2"
)

var bot *telebot.Bot
var msg *telebot.Message
var sendToGroup, sendToPrivate, sendTo telebot.Recipient
var textMsg, firstName, lastName, botName string
var pattern, typeChat, contentMessage, baseCommand string
var userName, chatTitle string
var userID int
var GoToFunc Func

type Func func() string

func Actions(b *telebot.Bot, m *telebot.Message, bName, baseCom string, groupID telebot.Recipient) {
	bot = b
	msg = m
	textMsg = m.Text
	userName = m.Sender.Username
	firstName = m.Chat.FirstName
	lastName = m.Chat.LastName
	chatTitle = m.Chat.Title
	botName = bName
	userID = m.Sender.ID
	baseCommand = baseCom
	sendToGroup = groupID
	sendToPrivate = m.Sender

	Private()
	Group()

	bot.Send(sendTo, contentMessage, telebot.ModeHTML)
}
