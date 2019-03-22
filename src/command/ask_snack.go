package command

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bandung-telegram-bot-golang/src/database"
	"github.com/bandung-telegram-bot-golang/src/helper"
	"github.com/bandung-telegram-bot-golang/src/message"
	"gopkg.in/tucnak/telebot.v2"
)

var TokenSnack, UsernameSnack, FullnameSnack string
var ChatIDSnack, PrivateIDSnack string

type userIDSnack struct {
	IDPrivateSnack string
}

func (u *userIDSnack) Recipient() string {
	return fmt.Sprintf("%s", u.IDPrivateSnack)
}

func Greet(n telebot.Recipient) string {
	return fmt.Sprintf("%s", n.Recipient())
}

func MatchAskSnack() string {
	pattern := strings.HasPrefix(textMsg, helper.PrefixCommandAskSnack())

	if pattern {
		GoToFunc = AskSnack
	} else {
		return "not match"
	}

	return GoToFunc()
}

func AskSnack() string {
	var snack string

	snack = helper.CheckEmptyAskSnack(textMsg)

	database.AskSnack(userName)

	if snack != "" {
		contentMessage = message.AskSnack(userName, snack)

		file, err := os.Open("temp.go")
		helper.ErrorMessage(err)
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			sendToUser := &userIDSnack{scanner.Text()}
			fmt.Println(sendToUser)

			bot.Send(sendToUser, contentMessage, telebot.ModeHTML)
		}
	} else {
		contentMessage = message.EmptyAskSnack(userName, baseCommand)
		sendTo = sendToPrivate
	}

	return "success"
}
