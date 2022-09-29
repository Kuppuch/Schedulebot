package main

import (
	"Schedulebot/pkg/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"time"
)

func main() {
	err := database.Connect()
	if err != nil {
		log.Fatalf("%v", err)
	}
	bot, err := tgbotapi.NewBotAPI("5528467830:AAH-axDRKLG25ECsDF43XGQx5BmdG5nMm1g")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	for {
		lessons := database.GetCurrentLessons()
		if len(lessons) > 0 {
			for _, v := range lessons {
				text := "Пара начнётся в течении 15 минут. Пара: " + v.Name + " Ссылка: " + v.Source
				msg := tgbotapi.NewMessage(538632285, text) // группа -599240202 / я 538632285
				bot.Send(msg)
			}
			time.Sleep(15 * time.Minute)
			continue
		}
		time.Sleep(time.Minute)
	}
}
