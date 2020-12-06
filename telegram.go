package main

import (
	//"fmt"
	"math/rand"
	"strconv"
	//"strings"
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

type Telegram struct {
	bot       *tgbotapi.BotAPI
	botConfig tgbotapi.UpdateConfig
}

func (t *Telegram) CreateBot() (err error) {
	t.bot, err = tgbotapi.NewBotAPI("1241791463:AAGTnqHu_2CMhPFAYTBloCr0tgriOTCHt0M")
	if err != nil {
		return err
	}
	t.botConfig = tgbotapi.NewUpdate(0)
	t.botConfig.Timeout = 60
	return nil
}

func (t Telegram) SendMessage(chatID int64, input string) error {
	_, err := t.bot.Send(tgbotapi.NewMessage(chatID, input))
	return err
}

func (t Telegram) SendReplyKeyboard(chatID int64) error {
	msg := tgbotapi.NewMessage(chatID, "Чтобы было проще хихикать, пользуйся клавиатурой.")
	msg.ReplyMarkup = replyKeyboard
	_, err := t.bot.Send(msg)
	return err
}

func (t Telegram) SendAnek(chatID int64, id int) error {
	if id < 0 || id > len(database.arrayOfAneks) {
		return nil
	}

	var likesKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(strconv.Itoa(database.Get(id).Likes)+" 👍🏻", "l"+strconv.Itoa(id)),
			tgbotapi.NewInlineKeyboardButtonData(strconv.Itoa(database.Get(id).Dislikes)+" 👎🏾", "d"+strconv.Itoa(id)),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🐖💨🤎", "f"+strconv.Itoa(id)),
		),
	)

	msg := tgbotapi.NewMessage(chatID, database.Get(id).Text)
	msg.ReplyMarkup = likesKeyboard
	_, err := t.bot.Send(msg)
	return err
}

func (t *Telegram) GetResponseFromInline(chatID int64, input string, callbackQuerryID string) {
	temp, _ := strconv.Atoi(input[1:len(input)])
	switch os := input[0]; os {
	case 'l':
		if !database.IsLike(chatID, temp) {
			database.Like(chatID, temp)
		} else {
			t.bot.AnswerCallbackQuery(tgbotapi.NewCallback(callbackQuerryID, "Ну ты и шалун! Любишь шалить! Лайк то ты уже поставил."))
		}
	case 'd':
		if !database.IsDislike(chatID, temp) {
			database.Dislike(chatID, temp)
		} else {
			t.bot.AnswerCallbackQuery(tgbotapi.NewCallback(callbackQuerryID, "Наверное у тебя сахар повышен, раз тебя так разозлил этот анек. Дизлайк уже стоял."))
		}
	case 'f':
		if !database.IsFavourite(chatID, temp) {
			database.AddToFavourites(chatID, temp)
		} else {
			t.bot.AnswerCallbackQuery(tgbotapi.NewCallback(callbackQuerryID, "Анек не смешной, а ты его второй раз в избранное добавляешь."))
		}
	}
}

func (t Telegram) CheckUpdates() error {
	updates, err := t.bot.GetUpdatesChan(t.botConfig)
	if err != nil {
		return err
	}

	for update := range updates {
		if update.CallbackQuery != nil {
			t.bot.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID, "Молодец! Твой палец записан, куда надо."))
			//t.bot.Send(tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data))
			t.GetResponseFromInline(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data, update.CallbackQuery.ID)
		}

		if update.Message == nil {
			continue
		}
		t.CreateAnswer(*update.Message)
	}
	return nil
}

func (t Telegram) CreateAnswer(input tgbotapi.Message) error {
	i, err := strconv.Atoi(input.Text)

	switch input.Text {
	
	case "/start":
		t.SendReplyKeyboard(input.Chat.ID)
	
	case "СЛУЧАЙНЫЙ АНЕК":
		t.SendAnek(input.Chat.ID, rand.Intn(anekQuantity))
	
	case "СЛУЧАЙНЫЙ СМЕШНОЙ АНЕК":
		_, index := database.GetLikedAnek()
		if index == 0 {
			t.SendMessage(input.Chat.ID, "Смешных анеков нет. Можешь посмотреть в зеркало.")
		}else{
			t.SendAnek(input.Chat.ID, index)
		}

	case "СЛУЧАЙНЫЙ НЕСМЕШНОЙ АНЕК":
		_, index := database.GetDislikedAnek()
		if index == 0 {
			t.SendMessage(input.Chat.ID, "Несмешных анеков нет. Смейся, любитель похохотать.")
		}else{
			t.SendAnek(input.Chat.ID, index)
		}

	case "СЛУЧАЙНЫЙ ИЗБРАННЫЙ АНЕК":
		

	case "СПИСОК ИЗБРАННЫХ АНЕКОВ":
		t.SendMessage(input.Chat.ID, database.GetStringOfFavourites(input.Chat.ID))
	
	default :
		t.SendMessage(input.Chat.ID, "Ты что, дурачок? Нажимай на кнопки, либо пиши число.")
	}

	if err == nil {
		t.SendAnek(input.Chat.ID, i)
	}

	return err
}
