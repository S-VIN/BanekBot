package main

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/gocolly/colly"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func getAnec(index int) string {
	var output string
	c := colly.NewCollector()

	c.OnHTML("article", func(e *colly.HTMLElement) {
		output = e.Text
	})

	c.Visit("https://baneks.ru/" + strconv.Itoa(index))

	return output
}

func main() {
	t0 := time.Now()

	bot, err := tgbotapi.NewBotAPI("1356963581:AAGPlUyAkofdhcehODZ-jvIv9Qu9T196pRQ")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		var number, err = strconv.Atoi(msg.Text)
		if err == nil {
			if number < 1 || number > 1141 {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Такого анекдота нет. Есть только 1-1142. Случайный Анекдот:"))
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, getAnec(rand.Intn(1142))))
			} else {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, getAnec(number)))
			}
		} else {
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Такого анекдота нет. Есть только 1-1142. Случайный Анекдот:"))
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, getAnec(rand.Intn(1142))))
		}

	}

	t1 := time.Now()
	fmt.Println(t1.Sub(t0))

}
