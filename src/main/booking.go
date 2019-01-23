package main

import (
	"chat_type/group"
	"chat_type/private"
	"helper"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/joho/godotenv"
	tb "gopkg.in/tucnak/telebot.v2"
)

var BotToken, BotName string

func init() {
	env := godotenv.Load()

	if env != nil {
		log.Fatal("Error loading .env file")
	}

	BotToken = os.Getenv("TOKEN_BOOKING")
	BotName = os.Getenv("BOT_BOOKING")
}

func main() {
	bot, err := tb.NewBot(tb.Settings{
		Token:  BotToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	bot.Handle(tb.OnText, func(m *tb.Message) {
<<<<<<< HEAD
		var msg [6]string
		var id [1]int

		msg[0] = m.Text
		msg[1] = m.Sender.Username
		msg[2] = m.Chat.FirstName
		msg[3] = m.Chat.LastName
		msg[4] = m.Chat.Title
		msg[5] = BotName
		id[0] = m.Sender.ID

		commandPrivate := []string{"/help", "/start"}
		commandGroup := []string{"/status_staging", "/add_staging", "/booking", "/done", "/add_oncall", "/oncall"}

		baseCommand := regexp.MustCompile(helper.RegexCompileBaseCommand()).FindString(m.Text)

		if !m.Private() {
			for i := 0; i < len(commandGroup); i++ {
				if baseCommand == commandGroup[i] {
					bot.Send(m.Chat, group.SendMessage(msg[0], msg[1], msg[4], msg[5], id[0]), tb.ModeHTML)
					bot.Delete(m)
				}
			}
		} else {
			for i := 0; i < len(commandPrivate); i++ {
				if baseCommand == commandPrivate[i] {
					bot.Send(m.Sender, private.SendMessage(msg[0], msg[2], msg[3], msg[5], id[0]), tb.ModeHTML)
				}
			}
=======
		if !m.Private() {
			bot.Send(m.Chat, group.SendMessage(m.Text, m.Sender.Username), tb.ModeHTML)
		} else {
			bot.Send(m.Sender, private.SendMessage(m.Text, m.Chat.FirstName, m.Chat.LastName), tb.ModeHTML)
>>>>>>> 94bcef742cbe47c008f1b4ba6cd7132bb3b99c64
		}
	})

	bot.Start()
}
