package main

import (
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/bandung-telegram-bot-golang/src/command"
	"github.com/bandung-telegram-bot-golang/src/helper"
	"github.com/bandung-telegram-bot-golang/src/message"
	"github.com/bandung-telegram-bot-golang/src/user"
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

	bot.Send(greetingBotOwner, message.OnlineMessage(FullnameSnack), tb.ModeHTML)
	bot.Send(postToGroup, message.OnlineMessage(FullnameSnack), tb.ModeHTML)

	bot.Handle(tb.OnText, func(m *tb.Message) {
		commandPrivate := []string{"/help", "/start", "/done"}
		commandGroup := []string{"/add_snack", "/move", "/permanent", "/delete", "/cancel", "/holiday", "/add_admin", "/delete_admin", "/list_admin", "/plat"}

		baseCommand := regexp.MustCompile(helper.RegexCompileBaseCommand()).FindString(m.Text)
		spammer := user.IsSpammer(m.Sender.Username, baseCommand)

		if !m.Private() {
			for i := 0; i < len(commandGroup); i++ {
				if baseCommand == commandGroup[i] {
					if spammer {
						bot.Send(m.Chat, message.UserSpammer(m.Sender.Username), tb.ModeHTML)
					} else {
						command.Actions(bot, m, UsernameSnack, baseCommand, postToGroup)
					}

					user.SaveSpammer(m.Sender.Username, baseCommand)
				}
			}
		} else {
			for i := 0; i < len(commandPrivate); i++ {
				if baseCommand == commandPrivate[i] {
					command.Actions(bot, m, UsernameSnack, baseCommand, postToGroup)
				}
			}
		}

		bot.Delete(m)
	})

	bot.Start()
}
