package kkdai

import (
	"github.com/google/uuid"
	ytDownload "github.com/kkdai/youtube/v2"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
	"io"
	"log"
	"os"
	"tg-music-bot/src/music"
)

type KdaiYTMusic struct{}

func (y *KdaiYTMusic) DownloadById(id string) (music.FileName, error) {

	client := ytDownload.Client{}
	video, err := client.GetVideo(id)
	if err != nil {
		return "", err
	}
	format := video.Formats.WithAudioChannels().FindByQuality("medium")
	stream, _, err := client.GetStream(video, format)
	if err != nil {
		return "", err
	}
	videoName, _ := uuid.NewUUID()

	file, err := os.Create(videoName.String() + ".mp4")
	if err != nil {
		return "", err
	}
	defer file.Close()
	_, err = io.Copy(file, stream)
	if err != nil {
		return "", err
	}

	err = ffmpeg_go.Input(videoName.String()+".mp4").Audio().OverWriteOutput().Output(
		videoName.String()+".mp3", ffmpeg_go.KwArgs{"ac": 2}).Run()

	log.Println(err)

	musicFileName := music.FileName(videoName.String() + ".mp3")
	return musicFileName, nil
}
