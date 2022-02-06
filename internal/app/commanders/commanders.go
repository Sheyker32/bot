package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Commander struct {
	Bot *tgbotapi.BotAPI
}

func NewCommander() *Commander {
	bot, err := tgbotapi.NewBotAPI("5268324809:AAGU0tZ-LNK55jdFmZN5iIaXGDujf78FGeA")

	if err != nil {
		log.Println(err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)
	return &Commander{bot}
}
func (c *Commander) Help(id int64) {
	c.Bot.Send(tgbotapi.NewMessage(id, "Это команда хелп"))
}
