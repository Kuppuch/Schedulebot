package main

import (
	"Schedulebot/pkg/database"
	"flag"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"time"
)

func main() {
	flag.Parse()
	err := database.Connect()
	if err != nil {
		log.Fatalf("%v", err)
	}
	bot, err := tgbotapi.NewBotAPI("5528467830:AAH-axDRKLG25ECsDF43XGQx5BmdG5nMm1g") // прод
	if err != nil {
		log.Panic(err)
	}

	//var chatID int64 = -1001811852540 // группа -1001811852540 / я 538632285
	chatID := *database.NFlag

	bot.Debug = true

	//log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	msg := tgbotapi.NewMessage(chatID, "Бот создан при поддержке https://vk.com/vlsu_schedule")
	bot.Send(msg)

	go func() {
		fl := true
		for {
			now := time.Now()

			if now.Hour() == 9 && fl {
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

				bot.Send(msg)
				fl = false
			} else if now.Hour() != 9 && !fl {
				fl = true
			}

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
			//time.Sleep(time.Minute)
			time.Sleep(time.Second)
		}
	}()
	updates := bot.GetUpdatesChan(u)

	date := 0
	for update := range updates {
		if date == update.Message.Date {
			continue
		} else {
			date = update.Message.Date
		}
		if update.Message != nil {

			switch update.Message.Text {
			case "/today":
				sendToday(bot, chatID, update)
			case "/today@Schedbotbot":
				sendToday(bot, chatID, update)
			default:
				sayAnything(bot, chatID, update)
				break
			}
		}
	}
}

func sendToday(bot *tgbotapi.BotAPI, chatID int64, update tgbotapi.Update) {
	lessons := database.GetToday()
	text := ""

	for _, v := range lessons {
		v.Start = v.Start.Add(-3 * time.Hour)
		t := string(v.Start.AppendFormat([]byte(""), "15:04"))
		text += fmt.Sprintf("Пара начнётся в " + t + "\n" +
			"Пара: " + v.Name + "\n " +
			"Ссылка: " + v.Source + "\n \n")
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ReplyMarkup = tgbotapi.ReplyKeyboardRemove{RemoveKeyboard: true}
	bot.Send(msg)
}

func sayAnything(bot *tgbotapi.BotAPI, chatID int64, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(chatID, update.Message.Text)
	msg.ReplyMarkup = tgbotapi.ReplyKeyboardRemove{RemoveKeyboard: true}
	bot.Send(msg)
}
