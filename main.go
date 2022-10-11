package main

import (
	"Schedulebot/pkg/database"
	"fmt"
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

	var chatID int64 = 538632285 // группа -599240202 / я 538632285

	bot.Debug = true

	//log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	msg := tgbotapi.NewMessage(chatID, "Бот создан при поддержке https://vk.com/vlsu_schedule")
	bot.Send(msg)

	go func() {
		for {
			lessons := database.GetCurrentLessons()
			if len(lessons) > 0 {
				for _, v := range lessons {
					text := fmt.Sprintf("Пара начнётся в течение <b>15 минут</b>. \n" +
						"Пара: " + v.Name + "\n " +
						"Ссылка: " + v.Source)
					msg := tgbotapi.NewMessage(chatID, text)
					msg.ParseMode = "html"
					bot.Send(msg)
				}
				time.Sleep(15 * time.Minute)
				continue
			}
			time.Sleep(time.Minute)
		}
	}()

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			if update.Message.Text != "Расписание на сегодня" {
				continue
			}
			var numericKeyboard = tgbotapi.NewReplyKeyboard(
				tgbotapi.NewKeyboardButtonRow(
					tgbotapi.NewKeyboardButton("Расписание на сегодня"),
				),
			)

			lessons := database.GetToday()
			text := ""

			for _, v := range lessons {
				v.Start = v.Start.Add(-3 * time.Hour)
				t := string(v.Start.AppendFormat([]byte(""), "15:04"))
				if err != nil {
					log.Println(err)
				}
				text += fmt.Sprintf("Пара начнётся в " + t + "\n" +
					"Пара: " + v.Name + "\n " +
					"Ссылка: " + v.Source + "\n \n")
			}

			msg = tgbotapi.NewMessage(chatID, text)
			msg.ReplyMarkup = numericKeyboard

			bot.Send(msg)
		}
	}
}
