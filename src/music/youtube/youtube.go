package youtube

import (
	"context"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"log"
	"os"
	"tg-music-bot/src/music"
)

func getTitleAndId(query string) (YTVideoInfo, error) {
	var videosInfo YTVideoInfo
	result, err := getVideosInfo(query)
	if err != nil {
		return videosInfo, err
	}
	log.Println(result)
	videosInfo.Id = result.Id.VideoId
	videosInfo.Title = result.Snippet.Title

	return videosInfo, nil
}

func getVideosInfo(query string) (*youtube.SearchResult, error) {
	var result *youtube.SearchResult
	ctx := context.Background()
	ytService, err := youtube.NewService(ctx, option.WithAPIKey(os.Getenv("YOUTUBE_API_TOKEN")))
	if err != nil {
		return result, err
	}

	parts := []string{"snippet"}
	response, err := ytService.Search.List(parts).Q(query).Do()
	if err != nil {
		return result, err
	}

	result = response.Items[0]

	return result, nil
}

type YTVideoInfo struct {
	Id    string
	Title string
	Music music.FileName
}

type MusicFileName string

type YTDownloader interface {
	DownloadById(id string) (music.FileName, error)
}

type CustomYTDownloader struct {
	BE YTDownloader
}

func (y *CustomYTDownloader) Download(query string) (YTVideoInfo, error) {
	var ytvideo YTVideoInfo
	video, err := getTitleAndId(query)
	if err != nil {
		return ytvideo, err
	}

	filename, err := y.BE.DownloadById(video.Id)
	if err != nil {
		return ytvideo, err
	}
	ytvideo.Id = video.Id
	ytvideo.Title = video.Title
	ytvideo.Music = filename

	return ytvideo, nil
}
