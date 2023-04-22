package main

import (
	"os"
	"tg-music-bot/config"
	"tg-music-bot/src/bot"
)

func main() {
	config.LoadConfig()
	f := bot.Fetcher{}
	b := bot.NewBot(os.Getenv("BOT_TOKEN"), f)

	b.Serve()
}
