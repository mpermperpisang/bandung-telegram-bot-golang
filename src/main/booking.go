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

	"github.com/joho/godotenv"
	tb "gopkg.in/tucnak/telebot.v2"
)

var BotToken, BotUsername, BotFullname, PrivateID string

type Private struct {
	IDPrivate string
}

func (u *Private) Recipient() string {
	return fmt.Sprintf("%s", u.IDPrivate)
}

func Greet(n tb.Recipient) string {
	return fmt.Sprintf("%s", n.Recipient())
}

func init() {
	env := godotenv.Load()
	helper.ErrorMessage(env)

	BotToken = os.Getenv("TOKEN_BOOKING")
	BotUsername = os.Getenv("BOT_BOOKING")
	BotFullname = os.Getenv("NAME_BOOKING")
	PrivateID = os.Getenv("ID_PRIVATE")
}

func main() {
	greetingBotOwner := &Private{PrivateID}
	bot, err := tb.NewBot(tb.Settings{
		Token:  BotToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	helper.ErrorMessage(err)

	bot.Send(greetingBotOwner, message.OnlineMessage(BotFullname), tb.ModeHTML)

	bot.Handle(tb.OnText, func(m *tb.Message) {
		var msg [7]string
		var id [1]int

		commandPrivate := []string{"/help", "/start", "/add_oncall"}
		commandGroup := []string{"/status_staging", "/add_staging", "/update_staging", "/booking", "/done", "/add_oncall", "/oncall"}

		baseCommand := regexp.MustCompile(helper.RegexCompileBaseCommand()).FindString(m.Text)

		msg[0] = m.Text
		msg[1] = m.Sender.Username
		msg[2] = m.Chat.FirstName
		msg[3] = m.Chat.LastName
		msg[4] = m.Chat.Title
		msg[5] = BotUsername
		msg[6] = baseCommand
		id[0] = m.Sender.ID

		if !m.Private() {
			for i := 0; i < len(commandGroup); i++ {
				if baseCommand == commandGroup[i] {
					bot.Send(m.Chat, group.SendMessage(msg[0], msg[1], msg[4], msg[5], msg[6], id[0]), tb.ModeHTML)
					bot.Delete(m)
				}
			}
		} else {
			for i := 0; i < len(commandPrivate); i++ {
				if baseCommand == commandPrivate[i] {
					bot.Send(m.Sender, private.SendMessage(msg[0], msg[2], msg[3], msg[5], msg[6], id[0]), tb.ModeHTML)
				}
			}
		}
	})

	bot.Start()
}
