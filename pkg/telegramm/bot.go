package telegramm

import (
	"DotaFind/pkg/repository"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot     *tgbotapi.BotAPI
	storage repository.ProfileRepository
}

func NewBot(bot *tgbotapi.BotAPI, pr repository.ProfileRepository) *Bot {
	return &Bot{bot: bot, storage: pr}
}

func (b *Bot) Start() {

	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	updates := b.initUpdatesChannel()
	b.handleUpdates(updates)
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {

	for update := range updates {

		if update.Message != nil {

			chat_id := update.Message.Chat.ID
			chat_id_BD, curent := b.storage.CheckUserInDB(chat_id)
			if chat_id == chat_id_BD.Id {

				if curent.Id >= 1 {

					if update.Message.Photo != nil {
						if curent.Id >= 2 {
							if curent.Id >= 3 {
								if curent.Id >= 4 {
									if update.Message.IsCommand() {
										b.handleCommand(update.Message)
										continue
									}
									b.handleMessage(update.Message)
								}
								msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Нельзя указывать возраст фоткой)!")
								b.bot.Send(msg)
							} else {
								msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Фото уже загружено введите имя!")
								b.bot.Send(msg)
							}

						} else {
							photo := update.Message.Photo[0]
							if err := b.storage.SavePhoto(chat_id, photo.FileID); err != nil {
								log.Fatal("Не удалось добавить фото в БД", err)
							}

							if err := b.storage.SwitchCurrent(chat_id, 2); err != nil {
								log.Fatal("Не удалось сменить CURRENT", err)
							}

							msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Фото принято, теперь отправьте свое имя(можно со стикерами)")
							b.bot.Send(msg)
						}

					} else {
						if curent.Id >= 2 {
							if curent.Id >= 3 {
								if curent.Id >= 4 {
									if update.Message.IsCommand() {
										b.handleCommand(update.Message)
										continue
									}
									b.handleMessage(update.Message)
								} else {
									age := update.Message.Text

									if _, err := strconv.Atoi(age); err != nil {

										log.Println(err)
										msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вводить можно только число!)")
										b.bot.Send(msg)

									} else {

										if err := b.storage.SaveAge(chat_id, age); err != nil {
											log.Fatal("Не удалось сохранить AGE", err)
										}

										if err := b.storage.SwitchCurrent(chat_id, 4); err != nil {
											log.Fatal("Не удалось сменить CURRENT", err)
										}

										msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ваш профиль готов поздравляю! Нажми на -> /help <- для помощи!")
										b.bot.Send(msg)
									}
								}

							} else {
								name := update.Message.Text
								if err := b.storage.SaveName(chat_id, name); err != nil {
									log.Fatal("Не удалось сохранить имя в БД", err)
								}
								if err := b.storage.SwitchCurrent(chat_id, 3); err != nil {
									log.Fatal("Не удалось сменить CURRENT", err)
								}
								msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Теперь введите ваш возраст")
								b.bot.Send(msg)
							}
						} else {
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Это не фото!")
							b.bot.Send(msg)
						}
					}
				}

			} else {
				current, err := b.storage.CreateProfile(chat_id)
				if err != nil {
					log.Fatal(err)
				}

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Отправьте мне любое ваше фото")
				b.bot.Send(msg)

				if err != nil {
					log.Println("\n Не удалось отправить сообщение пользователю", err)
				}
				log.Println("\n Профиль создан,состояние: ", current)
			}
		}
	}
}

func (b *Bot) initUpdatesChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}
