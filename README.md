<h1>Go Music Bot</h1>

Simple Telegram bot for download YouTube Video, convert it to .mp3 music format and send to Telegram chat.

This bot uses library *https://github.com/kkdai/youtube* for download videos and *https://github.com/u2takey/ffmpeg-go* for converting audio.

<h2>Usage:</h2>
1. Create config .env file `cp ./config/.example.env ./config/.env` and insert yours BOT_TOKEN and YOUTUBE_API_TOKEN
2. Build binary file `go build cmd/bot/main.go`
3. Run bot `./main`