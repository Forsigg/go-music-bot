package bot

import (
	"log"
	"os"
	"strings"
	"tg-music-bot/src/music/youtube/kkdai"
)
import "tg-music-bot/src/music/youtube"

func logErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func isCommandName(command string, msgText string) bool {
	return strings.Contains(msgText, command)
}

func handleCommand(upd Update, b *Bot) {
	if isCommandName("/track", upd.Message.Text) {
		err := commandTrack(upd, b)
		logErr(err)
	} else if isCommandName("/help", upd.Message.Text) || isCommandName("/start", upd.Message.Text) {
		commandHelp(upd, b)
	}
}

func handleMessage(upd Update, b *Bot) {
	b.sendMessage(upd.Message, helpMessage)
}

func commandTrack(upd Update, b *Bot) error {
	y := youtube.CustomYTDownloader{BE: &kkdai.KdaiYTMusic{}}
	songName := strings.Trim(upd.Message.Text, "/track ")

	if strings.TrimSpace(songName) == "" {
		b.sendMessage(upd.Message, emptyTrackMessage)
		return nil
	}

	b.sendMessage(upd.Message, waitingMessage)
	video, err := y.Download(songName)
	if err != nil {
		b.sendMessage(upd.Message, errorMessage)
		return err
	}

	b.sendAudio(upd.Message, video.Music, video.Title)
	err = os.Remove(string(video.Music))
	err = os.Remove(strings.ReplaceAll(string(video.Music), ".mp3", ".mp4"))

	return nil
}

func commandHelp(upd Update, b *Bot) {
	b.sendMessage(upd.Message, helpMessage)
}
