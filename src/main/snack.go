package main

import (
	"chat_type/group"
	"chat_type/private"
	"fmt"
	"helper"
	"message"
	"os"
	"regexp"
	"time"
	"user"

	"github.com/joho/godotenv"
	tb "gopkg.in/tucnak/telebot.v2"
)

var TokenSnack, UsernameSnack, FullnameSnack string
var ChatIDSnack, PrivateIDSnack string

type PrivateSnack struct {
	IDPrivateSnack string
}

type GroupSnack struct {
	IDGroupSnack string
}

func (u *PrivateSnack) Recipient() string {
	return fmt.Sprintf("%s", u.IDPrivateSnack)
}

func (u *GroupSnack) Recipient() string {
	return fmt.Sprintf("%s", u.IDGroupSnack)
}

func Greet(n tb.Recipient) string {
	return fmt.Sprintf("%s", n.Recipient())
}

func init() {
	env := godotenv.Load()
	helper.ErrorMessage(env)

	//env untuk bot snack
	TokenSnack = os.Getenv("TOKEN_SNACK")
	UsernameSnack = os.Getenv("BOT_SNACK")
	FullnameSnack = os.Getenv("NAME_SNACK")
	ChatIDSnack = os.Getenv("ID_SNACK")
	//env untuk id bot owner
	PrivateIDSnack = os.Getenv("ID_PRIVATE")
}

func main() {
	greetingBotOwner := &PrivateSnack{PrivateIDSnack}
	postToGroup := &GroupSnack{ChatIDSnack}
	bot, err := tb.NewBot(tb.Settings{
		Token:  TokenSnack,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	helper.ErrorMessage(err)

	fmt.Println(postToGroup)

	bot.Send(greetingBotOwner, message.OnlineMessage(FullnameSnack), tb.ModeHTML)

	bot.Handle(tb.OnText, func(m *tb.Message) {
		var msg [7]string
		var id [1]int

		commandPrivate := []string{"/help", "/start", "/done"}
		commandGroup := []string{"/add_snack", "/move", "/permanent", "/delete", "/cancel", "/holiday", "/add_admin", "/delete_admin", "/list_admin", "/schedule"}

		baseCommand := regexp.MustCompile(helper.RegexCompileBaseCommand()).FindString(m.Text)
		spammer := user.IsSpammer(m.Sender.Username, baseCommand)

		msg[0] = m.Text
		msg[1] = m.Sender.Username
		msg[2] = m.Chat.FirstName
		msg[3] = m.Chat.LastName
		msg[4] = m.Chat.Title
		msg[5] = UsernameSnack
		msg[6] = baseCommand
		id[0] = m.Sender.ID

		if !m.Private() {
			for i := 0; i < len(commandGroup); i++ {
				if baseCommand == commandGroup[i] {
					if spammer {
						bot.Send(m.Chat, message.UserSpammer(m.Sender.Username), tb.ModeHTML)
					} else {
						bot.Send(m.Chat, group.SendMessage(msg[0], msg[1], msg[2], msg[3], msg[4], msg[5], msg[6], id[0]), tb.ModeHTML)
					}

					user.SaveSpammer(m.Sender.Username, baseCommand)
					bot.Delete(m)
				}
			}
		} else {
			for i := 0; i < len(commandPrivate); i++ {
				if baseCommand == commandPrivate[i] {
					bot.Send(m.Sender, private.SendMessage(msg[0], msg[1], msg[2], msg[3], msg[4], msg[5], msg[6], id[0]), tb.ModeHTML)
				}
			}
		}
	})

	bot.Start()
}
