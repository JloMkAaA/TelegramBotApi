package main

import (
	"DotaFind/pkg/repository/sqlite"
	"DotaFind/pkg/telegramm"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const patch = "data/sqlite.db"

func main() {

	bot, err := tgbotapi.NewBotAPI("5667089545:AAENcFYGOUdStzAN-IBYKPIcdfZSzAfOICI")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	db, err := sql.Open("sqlite3", patch)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	storage := sqlite.NewStorageRepository(db)
	if err != nil {
		log.Fatal("can't connect to DB", err)
	}

	if err := storage.Init(); err != nil {
		log.Fatal(err)
	}

	TelegrammBot := telegramm.NewBot(bot, storage)
	TelegrammBot.Start()
}
