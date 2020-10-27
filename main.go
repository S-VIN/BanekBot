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
		//fmt.Println(e.Text)
		output = e.Text
	})

	c.Visit("https://baneks.ru/" + strconv.Itoa(index))

	return output
}

func main() {
	t0 := time.Now()

	//for i := 1; i < 1143; i++ {
	//	fmt.Println(getAnec(i))
	//}

	bot, err := tgbotapi.NewBotAPI("1134594213:AAFJaUZZCGnFdRANSBIfgF0YJBn-VJS9nTc")
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

		//msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, getAnec(rand.Intn(1142))))
	}

	t1 := time.Now()
	fmt.Println(t1.Sub(t0))

}
