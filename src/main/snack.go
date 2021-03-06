package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/bandung-telegram-bot-golang/src/command"
	"github.com/bandung-telegram-bot-golang/src/database"
	"github.com/bandung-telegram-bot-golang/src/helper"
	"github.com/bandung-telegram-bot-golang/src/message"
	"github.com/bandung-telegram-bot-golang/src/user"
	"github.com/joho/godotenv"
	tb "gopkg.in/tucnak/telebot.v2"
)

var TokenSnack, UsernameSnack, FullnameSnack string
var ChatIDSnack, PrivateIDSnack, ChatGroupSnack string

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

	// env untuk bot snack
	TokenSnack = os.Getenv("TOKEN_SNACK")
	UsernameSnack = os.Getenv("BOT_SNACK")
	FullnameSnack = os.Getenv("NAME_SNACK")
	ChatIDSnack = os.Getenv("ID_SNACK")
	ChatGroupSnack = os.Getenv("ID_GROUP")
	// env untuk id bot owner
	PrivateIDSnack = os.Getenv("ID_PRIVATE")
}

func main() {
	greetingBotOwner := &PrivateSnack{PrivateIDSnack}
	postToGroup := &GroupSnack{ChatIDSnack}
	postToBdg := &GroupSnack{ChatGroupSnack}
	bot, err := tb.NewBot(tb.Settings{
		Token:  TokenSnack,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	helper.ErrorMessage(err)

	bot.Send(greetingBotOwner, message.OnlineMessage(FullnameSnack), tb.ModeHTML)
	// bot.Send(postToGroup, message.OnlineMessage(FullnameSnack), tb.ModeHTML)

	bot.Handle(tb.OnAddedToGroup, func(m *tb.Message) {
		bot.Send(m.Chat, message.AddedGroup(m.Chat.Title), tb.ModeHTML)
	})

	bot.Handle(tb.OnUserJoined, func(m *tb.Message) {
		fmt.Println(m.UserJoined.LanguageCode)
		// if m.UserJoined.LanguageCode != "" {
		bot.Send(m.Chat, message.UserJoin(m.Chat.Title, m.UserJoined.Username), tb.ModeHTML)
		database.AddOnboarding("@" + m.UserJoined.Username)
		bot.Send(m.UserJoined, message.OnboardingUser(m.UserJoined.FirstName, m.UserJoined.LastName), tb.ModeHTML)
		// }
	})

	bot.Handle(tb.OnUserLeft, func(m *tb.Message) {
		fmt.Println(m.UserLeft.LanguageCode)
		// if m.UserLeft.LanguageCode != "" {
		bot.Send(m.Chat, message.UserLeft("@"+m.UserLeft.Username), tb.ModeHTML)
		database.DeleteAdminSnack("@" + m.UserLeft.Username)
		database.DeleteSnack("@" + m.UserLeft.Username)
		database.DeleteOnboarding("@" + m.UserLeft.Username)
		bot.Send(m.UserLeft, message.UserLeft(m.UserLeft.FirstName), tb.ModeHTML)
		// }
	})

	bot.Handle(tb.OnText, func(m *tb.Message) {
		commandBot := []string{"/help", "/start", "/baik", "/done", "/add_snack", "/move", "/permanent", "/delete", "/cancel", "/holiday", "/add_admin", "/delete_admin", "/list_admin", "/ask"}
		commandGroup := []string{"/baik", "/add_snack", "/move", "/permanent", "/delete", "/cancel", "/holiday", "/add_admin", "/delete_admin", "/list_admin", "/ask"}
		baseCommand := regexp.MustCompile(helper.RegexCompileBaseCommand()).FindString(m.Text)
		spammer := user.IsSpammer(m.Sender.Username, baseCommand)
		includeCommand, _ := helper.IncludeArray(baseCommand, commandBot)

		if !m.Private() && !database.NewOnboarding(m.Sender.Username) {
			bot.Send(m.Chat, "Kak @"+m.Sender.Username+", please japri @"+os.Getenv("BOT_SNACK")+" dulu yaa.. Nuhun")
		} else if m.Private() && !database.NewOnboarding(m.Sender.Username) {
			content := command.NewOnboarding(m.Text, m.Sender.Username)
			bot.Send(m.Sender, content, tb.ModeHTML)

			if strings.Contains(content, "selanjutnya") {
				bot.Send(postToBdg, message.OnboardingSuccess(m.Text, m.Sender.Username), tb.ModeHTML)
			}
		} else {
			for i := 0; i < len(commandBot); i++ {
				if !m.Private() {
					if baseCommand == commandGroup[i] {
						if spammer {
							bot.Send(m.Sender, message.UserSpammer(m.Sender.Username), tb.ModeHTML)
						} else {
							command.Actions(bot, m, UsernameSnack, baseCommand, postToGroup)
						}

						user.SaveSpammer(m.Sender.Username, baseCommand)
					}

					if includeCommand {
						bot.Delete(m)
					}
				} else {
					if baseCommand == commandBot[i] {
						command.Actions(bot, m, UsernameSnack, baseCommand, postToGroup)
					}
				}
			}
		}
	})

	bot.Start()
}
