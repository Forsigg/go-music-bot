package bot

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"tg-music-bot/src/music/youtube/kkdai"
	types "tg-music-bot/src/telegram_types"
)
import "tg-music-bot/src/music/youtube"

func logErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func addUserInDBifNotExists(upd types.Update, b *Bot) {
	_, err := b.db.GetUserById(upd.Message.From.Id)
	if err != nil {
		err = b.db.AddUser(upd.Message.From)
		if err != nil {
			log.Printf("Error while adding user in database: %s", err)
		}
	}
}

func isCommandName(command string, msgText string) bool {
	return strings.Contains(msgText, command)
}

func handleCommand(upd types.Update, b *Bot) {
	if isCommandName("/track", upd.Message.Text) {
		err := commandTrack(upd, b)
		logErr(err)
	} else if isCommandName("/help", upd.Message.Text) || isCommandName("/start", upd.Message.Text) {
		commandHelp(upd, b)
	} else if isCommandName("/send", upd.Message.Text) {
		commandSend(upd, b)
	}
}

func handleMessage(upd types.Update, b *Bot) {
	b.sendMessage(upd.Message.Chat.Id, helpMessage)
}

func commandTrack(upd types.Update, b *Bot) error {
	y := youtube.CustomYTDownloader{BE: &kkdai.KdaiYTMusic{}}
	songName := strings.Trim(upd.Message.Text, "/track ")

	if strings.TrimSpace(songName) == "" {
		b.sendMessage(upd.Message.Chat.Id, emptyTrackMessage)
		return nil
	}

	b.sendMessage(upd.Message.Chat.Id, waitingMessage)
	video, err := y.Download(songName)
	if err != nil {
		b.sendMessage(upd.Message.Chat.Id, errorMessage)
		return err
	}

	b.sendAudio(upd.Message, video.Music, video.Title)
	err = os.Remove(string(video.Music))
	err = os.Remove(strings.ReplaceAll(string(video.Music), ".mp3", ".mp4"))

	return nil
}

func commandHelp(upd types.Update, b *Bot) {
	addUserInDBifNotExists(upd, b)
	b.sendMessage(upd.Message.Chat.Id, helpMessage)
}

func commandSend(upd types.Update, b *Bot) {
	adminId, _ := strconv.Atoi(os.Getenv("ADMIN_ID"))
	if upd.Message.From.Id != adminId {
		return
	}
	msgToSend := strings.Trim(upd.Message.Text, "/send ")
	users, err := b.db.GetAllUsers()
	if err != nil {
		b.sendMessage(upd.Message.Chat.Id, fmt.Sprintf("Error while start sending message to all users: %s", err))
		return
	}

	wg := sync.WaitGroup{}

	for _, user := range users {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			b.sendMessage(id, msgToSend)
		}(user.Id)
	}
	wg.Wait()
}
