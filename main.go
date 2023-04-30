package main

import (
	"log"
	"translator/commands"
	"translator/e"
	"translator/translate"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6030853445:AAFgSZvUEWa0u0WrgtlaYQFz47-iiIaCMOs")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {

		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			switch update.Message.Text {
			case "/start":
				bot.Send(commands.StartMsg(update))
				continue
			case "/help":
				bot.Send(commands.Help(update))
				continue
			}

			translatedText, err := translate.Translate(update.Message.Text)
			if err != nil {
				e.Wrap("can't translate message", err)
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, translatedText)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
