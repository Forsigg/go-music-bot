package main

import (
	"fmt"
	"log"
	"os"
	"tg-music-bot/config"
	"tg-music-bot/src/bot"
	"tg-music-bot/src/database/sqlite"
)

func main() {
	config.LoadConfig()
	db, err := sqlite.NewSQLiteDB(os.Getenv("DB_PATH"))
	fmt.Println(os.Getenv("DB_PATH"))
	if err != nil {
		log.Fatal(err)
	}
	f := bot.Fetcher{}
	b := bot.NewBot(os.Getenv("BOT_TOKEN"), f, db)

	b.Serve()
}
