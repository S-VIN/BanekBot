package main

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var telegram Telegram

var replyKeyboard = tgbotapi.NewReplyKeyboard(
    tgbotapi.NewKeyboardButtonRow(
        tgbotapi.NewKeyboardButton("СЛУЧАЙНЫЙ АНЕК"),
    ),
    tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("СЛУЧАЙНЫЙ СМЕШНОЙ АНЕК"),
		tgbotapi.NewKeyboardButton("СЛУЧАЙНЫЙ НЕСМЕШНОЙ АНЕК"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("СЛУЧАЙНЫЙ ИЗБРАННЫЙ АНЕК"),
		tgbotapi.NewKeyboardButton("СПИСОК ИЗБРАННЫХ АНЕКОВ"),
	),	
)

var likesKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("👍🏻","like"),
		tgbotapi.NewInlineKeyboardButtonData("👎🏾","dislike"),
	),
)

type Telegram struct{
	bot *tgbotapi.BotAPI
	botConfig tgbotapi.UpdateConfig
	
}


func (t *Telegram)CreateBot() (err error) {
	t.bot, err = tgbotapi.NewBotAPI("1241791463:AAGTnqHu_2CMhPFAYTBloCr0tgriOTCHt0M")
	if err != nil {
		return err
	}
	t.botConfig = tgbotapi.NewUpdate(0)
	t.botConfig.Timeout = 60
	return nil
}

func (t Telegram)SendMessage(chatID int64, input string) error {
	_, err := t.bot.Send(tgbotapi.NewMessage(chatID, input))
	return err
}

func (t Telegram)SendMessageWithReply(chatID int64, k tgbotapi.ReplyKeyboardMarkup, ) error {
	msg := tgbotapi.NewMessage(chatID, "Чтобы было проще хихикать, пользуйся клавиатурой.")
	msg.ReplyMarkup = k
	_, err := t.bot.Send(msg)
	return err
}


func (t Telegram)SendMessageWithInline(chatID int64) error {
	msg := tgbotapi.NewMessage(chatID, "Чтобы было проще хихикать, пользуйся клавиатурой.")
	msg.ReplyMarkup = inKeyboard
	_, err := t.bot.Send(msg)
	return err
}

func (t Telegram)CheckUpdates() error {
	updates, err := t.bot.GetUpdatesChan(t.botConfig)
	if err != nil {
		return err
	}
	for update := range updates {
		if update.Message == nil {
			continue
		}
		t.CreateAnswer(*update.Message)
	}
	return nil
}

func (t Telegram)CreateAnswer(input tgbotapi.Message) error {
	i, err := strconv.Atoi(input.Text)
	fmt.Println(input.Text)
	temp := database.Get(i)
	if input.Text == "/start"{
		t.SendKeyboard(input.Chat.ID, numericKeyboard)
		fmt.Println("key")
	}
	if input.Text == "t"{
		t.SendSpecial(input.Chat.ID)
		fmt.Println("key")
	}
	if err == nil {
		t.SendMessage(input.Chat.ID, temp.Text)
	}
	return err
}
