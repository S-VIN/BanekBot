package main

import (
	//"fmt"
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

func (t Telegram)SendReplyKeyboard(chatID int64) error {
	msg := tgbotapi.NewMessage(chatID, "Чтобы было проще хихикать, пользуйся клавиатурой.")
	msg.ReplyMarkup = replyKeyboard
	_, err := t.bot.Send(msg)
	return err
}

func (t Telegram)SendAnek(chatID int64, id int) error{
	if id < 0 || id > len(database.arrayOfAneks){
		return nil
	}
	
	var likesKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(strconv.Itoa(database.Get(id).Likes) + " 👍🏻",  "l" + strconv.Itoa(id)),
			tgbotapi.NewInlineKeyboardButtonData(strconv.Itoa(database.Get(id).Dislikes) + " 👎🏾", "d" + strconv.Itoa(id)),
		),
	)

	msg := tgbotapi.NewMessage(chatID, database.Get(id).Text)
	msg.ReplyMarkup = likesKeyboard
	_, err := t.bot.Send(msg)
	return err
}

func (t *Telegram)UpdateLikes(input string){
	var like bool
	if input[0] == 'l'{
		like = true
	} else{
		like = false
	}
	temp, _ := strconv.Atoi(input[1:len(input)])  
	if like{
		database.Like(temp)
	}else{
		database.Dislike(temp)
	}
}

func (t Telegram)CheckUpdates() error {
	updates, err := t.bot.GetUpdatesChan(t.botConfig)
	if err != nil {
		return err
	}
	
	for update := range updates {
		if update.CallbackQuery != nil{
			t.bot.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID, "Молодец!"))
			//t.bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data))
			t.UpdateLikes(update.CallbackQuery.Data)
		
		}

		if update.Message == nil {
			continue
		}
		t.CreateAnswer(*update.Message)
	}
	return nil
}

func (t Telegram)CreateAnswer(input tgbotapi.Message) error {
	i, err := strconv.Atoi(input.Text)
	
	if input.Text == "/start"{
		t.SendReplyKeyboard(input.Chat.ID)
	}
	
	if err == nil {
		t.SendAnek(input.Chat.ID, i)
	}
	
	return err
}
