package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"tg-music-bot/src/music"
	"time"
)

type Bot struct {
	token   string
	fetcher Fetcher
}

func NewBot(token string, fetcher Fetcher) *Bot {
	return &Bot{
		token:   token,
		fetcher: fetcher,
	}
}

func (b *Bot) processUpdate(upd Update) {
	if upd.Message.IsCommand() {
		handleCommand(upd, b)
	} else {
		handleMessage(upd, b)
	}
}

func (b *Bot) Serve() {
	log.Println("Bot start serving...")
	offset := -1
	for {
		updates := b.fetcher.Fetch(b.token, offset)
		if len(updates) == 0 {
			time.Sleep(2 * time.Second)
		}
		for _, update := range updates {
			fmt.Println(update.Id, update.Message.From.Username, update.Message.Text)
			go b.processUpdate(update)
			offset = update.Id + 1
		}
	}
}

func (b *Bot) sendAudio(msg Message, file music.FileName, trackName string) {
	u := b.buildUrl(sendAudioMethod)

	values := map[string]io.Reader{
		"chat_id": strings.NewReader(strconv.Itoa(msg.Chat.Id)),
		"audio":   mustOpen(string(file)),
		"caption": strings.NewReader(trackName),
	}

	err := uploadFile(u, values)
	log.Println(err)
}

func (b *Bot) sendMessage(msg Message, text string) {
	u := b.buildUrl(sendMessageMethod)
	values := map[string]string{
		"chat_id": strconv.Itoa(msg.Chat.Id),
		"text":    text,
	}

	jsonData, err := json.Marshal(values)
	if err != nil {
		log.Printf("Error while building json in SendMessage: %s", err)
	}
	_, err = http.Post(u, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Error while send http request to telegram API sendMessage: %s", err)
	}

}

func (b *Bot) buildUrl(method string) string {
	u, err := url.Parse(TgApiUrl + b.token + "/" + method)
	if err != nil {
		log.Printf("Error while parsing base tgAPI url: %s", err)
	}
	return u.String()
}

func uploadFile(url string, values map[string]io.Reader) (err error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	for key, r := range values {
		var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}

		if x, ok := r.(*os.File); ok {
			if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
				return
			}
		} else {
			if fw, err = w.CreateFormField(key); err != nil {
				return
			}
		}
		if _, err = io.Copy(fw, r); err != nil {
			return err
		}

	}
	w.Close()

	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return
	}

	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status: %s", res.Status)
	}
	return
}

func mustOpen(f string) *os.File {
	r, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	return r
}