package main

import (
	"log"

	commanders "github.com/Sheyker32/bot/internal/app/commanders"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	// подключаемся к боту с помощью токена
	Commander := commanders.NewCommander()

	// инициализируем канал, куда будут прилетать обновления от API
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := Commander.Bot.GetUpdatesChan(u)

	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			//msg.ReplyToMessageID = update.Message.MessageID
			switch update.Message.Command() {
			case "help":
				Commander.Help(update.Message.Chat.ID)

			default:
				//bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text))
			}

		}
	}

}
