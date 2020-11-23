package main

import 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

var Bot *BOTApi
var BotConfig tgbotapi.UpdateConfig


func CreateBot() (err error){
	Bot, err := tgbotapi.NewBotAPI("1241791463:AAGTnqHu_2CMhPFAYTBloCr0tgriOTCHt0M")
	if err != nil{
		return err
	}

	Bot.Debug = false
	
	BotUpdates := tgbotapi.NewUpdate(0)
	BotUpdates.Timeout = 60

	return nil
}


func CheckUpdates() error {
	 
	var updates[] tgbotapi.Update
	updates, err = tgbotapi.GetUpdates(BotConfig)
	for _, update := range updates{
		SendMessage(update.Message)
	}

}

func SendMessage(input string) error{
	
	Bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, input))

}




	if update.Message == nil { // ignore any non-Message Updates
		continue
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

	var number, err = strconv.Atoi(msg.Text)
	if err == nil {
		if number < -2 || number > 1142 {
			
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, getAnecFromDatabase(rand.Intn(1142), conn)))
		} else {
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, getAnecFromDatabase(number, conn)))
		}
	} else {
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Такого анекдота нет. Есть только 1-1142. Случайный Анекдот:"))
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, getAnecFromDatabase(rand.Intn(1142), conn)))
	}

