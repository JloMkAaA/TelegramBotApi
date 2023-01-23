package telegramm

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/start"),
		tgbotapi.NewKeyboardButton("/profile"),
	),
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	switch message.Command() {
	case "start":
		return b.handleStartCommand(message)
	case "profile":
		return b.handleProfileCommand(message)
	case "help":
		return b.handleHelpCommand(message, msg)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Я не знаю такой команды! Нажми на -> /help <- для помощи!")
	b.bot.Send(msg)

	return nil
}

func (b *Bot) handleHelpCommand(message *tgbotapi.Message, msg tgbotapi.MessageConfig) error {
	msg.ReplyMarkup = numericKeyboard
	b.bot.Send(msg)

	return nil
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Ты ввел команду /start")
	b.bot.Send(msg)

	return nil
}

func (b *Bot) handleProfileCommand(message *tgbotapi.Message) error {

	profile, err := b.storage.GetProfile(message.Chat.ID)

	if err != nil {
		log.Fatal("Не удалось получить профиль", err)
	}

	// data := tgbotapi.Read

	// file := tgbotapi.NewInputMediaPhoto(profile.Photo)

	str := fmt.Sprint(profile.Name, ", ", profile.Age)

	msg := tgbotapi.NewPhoto(message.Chat.ID, tgbotapi.FileID(profile.Photo))
	msg.Caption = str
	b.bot.Send(msg)

	return nil
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {

	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, "Это обычное сообщение")
	b.bot.Send(msg)

	return nil
}
