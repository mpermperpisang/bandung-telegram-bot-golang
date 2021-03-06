package main

import (
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/bandung-telegram-bot-golang/src/command"
	"github.com/bandung-telegram-bot-golang/src/helper"
	"github.com/bandung-telegram-bot-golang/src/message"
	"github.com/joho/godotenv"
	tb "gopkg.in/tucnak/telebot.v2"
)

var TokenBooking, UsernameBooking, FullnameBooking string
var ChatIDBooking, PrivateIDBooking string

type PrivateBooking struct {
	IDPrivateBooking string
}

type GroupBooking struct {
	IDGroupBooking string
}

func (u *PrivateBooking) Recipient() string {
	return fmt.Sprintf("%s", u.IDPrivateBooking)
}

func (u *GroupBooking) Recipient() string {
	return fmt.Sprintf("%s", u.IDGroupBooking)
}

func GreetBooking(n tb.Recipient) string {
	return fmt.Sprintf("%s", n.Recipient())
}

func init() {
	env := godotenv.Load()
	helper.ErrorMessage(env)

	// env untuk bot booking
	TokenBooking = os.Getenv("TOKEN_BOOKING")
	UsernameBooking = os.Getenv("BOT_BOOKING")
	FullnameBooking = os.Getenv("NAME_BOOKING")
	ChatIDBooking = os.Getenv("ID_BOOKING")
	// env untuk id bot owner
	PrivateIDBooking = os.Getenv("ID_PRIVATE")
}

func main() {
	greetingBotOwner := &PrivateBooking{PrivateIDBooking}
	postToGroup := &GroupBooking{ChatIDBooking}
	bot, err := tb.NewBot(tb.Settings{
		Token:  TokenBooking,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	helper.ErrorMessage(err)

	bot.Send(greetingBotOwner, message.OnlineMessage(FullnameBooking), tb.ModeHTML)
	// bot.Send(postToGroup, message.OnlineMessage(FullnameBooking), tb.ModeHTML)

	bot.Handle(tb.OnAddedToGroup, func(m *tb.Message) {
		bot.Send(m.Chat, message.AddedGroup(m.Chat.Title), tb.ModeHTML)
	})

	bot.Handle(tb.OnText, func(m *tb.Message) {
		commandBot := []string{"/help", "/start", "/status_staging", "/add_staging", "/update_staging", "/book_staging", "/done_staging", "/add_oncall", "/oncall"}
		commandGroup := []string{"/status_staging", "/add_staging", "/update_staging", "/book_staging", "/done_staging", "/add_oncall", "/oncall"}
		baseCommand := regexp.MustCompile(helper.RegexCompileBaseCommand()).FindString(m.Text)
		includeCommand, _ := helper.IncludeArray(baseCommand, commandBot)

		for i := 0; i < len(commandBot); i++ {
			if !m.Private() {
				if baseCommand == commandGroup[i] {
					command.Actions(bot, m, UsernameBooking, baseCommand, postToGroup)
				}

				if includeCommand {
					bot.Delete(m)
				}
			} else {
				if baseCommand == commandBot[i] {
					command.Actions(bot, m, UsernameBooking, baseCommand, postToGroup)
				}
			}
		}
	})

	bot.Start()
}
