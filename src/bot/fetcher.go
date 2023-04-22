package bot

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

const TgApiUrl string = "https://api.telegram.org/bot"

type Fetcher struct {
}

func (f *Fetcher) Fetch(token string, offset int) []Update {
	return f.ParseUpdates(f.makeRequest(token, offset))
}

func (f *Fetcher) makeRequest(token string, offset int) []byte {
	u := buildGetUpdatesUrl(token, offset)
	resp, err := http.Get(u)
	if err != nil {
		log.Printf("Error while send request to %s: %s", u, err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error while reading response body from %s: %s", u, err)
	}
	return body
}

func (f *Fetcher) ParseUpdates(respBody []byte) []Update {
	var result ResponseResult
	err := json.Unmarshal(respBody, &result)
	if err != nil {
		log.Printf("Error while parsing json: %s", err)
		return nil
	}
	return result.Updates
}

func buildGetUpdatesUrl(token string, offset int) string {
	u, err := url.Parse(TgApiUrl + token + "/" + getUpdatesMethod)
	if err != nil {
		log.Printf("Error while parsing tgAPI url: %s", err)
	}
	q := u.Query()
	q.Set("offset", strconv.Itoa(offset))
	u.RawQuery = q.Encode()
	return u.String()
}
