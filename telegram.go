package main

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var Bot *tgbotapi.BotAPI
var BotConfig tgbotapi.UpdateConfig


func CreateBot() (err error){
	Bot, err := tgbotapi.NewBotAPI("1241791463:AAGTnqHu_2CMhPFAYTBloCr0tgriOTCHt0M")
	if err != nil{
		return err
	}

	Bot.Debug = true
	
	BotUpdates := tgbotapi.NewUpdate(0)
	BotUpdates.Timeout = 60

	return nil
}


func CheckUpdates() error {
	var updates[] tgbotapi.Update
	var err error
	updates, err = Bot.GetUpdatesChan(BotConfig)
	if err != nil{
		return err
	}
	for update := range updates{
		if update.Message == nil{
			continue
		}
		SendMessage(update.Message.Chat.ID, update.Message.Text)
	}
	return nil
}

func SendMessage(chatID int64, input string) error{
	mes, err := Bot.Send(tgbotapi.NewMessage(chatID, input))
	fmt.Println(mes)
	return err
}
