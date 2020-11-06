package main

import (
	"context"
	"math/rand"
	"strconv"

	//"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/gocolly/colly"
	"github.com/jackc/pgx"
)

func getAnecFromInternet(index int) string {
	var output string
	c := colly.NewCollector()

	c.OnHTML("article", func(e *colly.HTMLElement) {
		output = e.Text
	})

	c.Visit("https://baneks.ru/" + strconv.Itoa(index))

	return output
}

func getAnecFromDatabase(index int, conn *pgx.Conn) string {
	var number int
	var anek string
	//_, err = conn.Exec(context.Background(), "INSERT INTO aneks(number, anek)VALUES ("+ strconv.Itoa(i) + ", '" + text + "');" )
	err := conn.QueryRow(context.Background(), "select number, anek from aneks where number="+strconv.Itoa(index)+";").Scan(&number, &anek)
	if err != nil {

	}
	return anek
}

func main() {

	conn, _ := pgx.Connect(context.Background(), "user=stepan password= host=localhost port=5432 dbname=stepan sslmode=verify-ca")
	defer conn.Close(context.Background())
	bot, _ := tgbotapi.NewBotAPI("1134594213:AAFJaUZZCGnFdRANSBIfgF0YJBn-VJS9nTc")
	bot.Debug = false
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	//fmt.Println(getAnecFromDatabase(50, conn))

	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		var number, err = strconv.Atoi(msg.Text)
		if err == nil {
			if number < -2 || number > 1142 {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Такого анекдота нет. Есть только 1-1142. Случайный Анекдот:"))
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, getAnecFromDatabase(rand.Intn(1142), conn)))
			} else {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, getAnecFromDatabase(number, conn)))
			}
		} else {
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Такого анекдота нет. Есть только 1-1142. Случайный Анекдот:"))
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, getAnecFromDatabase(rand.Intn(1142), conn)))
		}

	}
}
