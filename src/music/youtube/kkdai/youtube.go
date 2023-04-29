package kkdai

import (
	"github.com/google/uuid"
	ytDownload "github.com/kkdai/youtube/v2"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
	"io"
	"log"
	"os"
	"tg-music-bot/config"
	"tg-music-bot/src/music"
)

type KdaiYTMusic struct{}

// DownloadById implements YTDownload interface method. Download video by id in arg,
// create file with uuid name, convert it in mp3 and return filename
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
	videoFileName := config.GetWorkdirFromOsEnv() + "/" + videoName.String() + ".mp4"
	musicFileName := music.FileName(config.GetWorkdirFromOsEnv() + "/" + videoName.String() + ".mp3")

	file, err := os.Create(videoFileName)
	if err != nil {
		return "", err
	}
	defer file.Close()
	_, err = io.Copy(file, stream)
	if err != nil {
		return "", err
	}

	err = ffmpeg_go.Input(videoFileName).Audio().OverWriteOutput().Output(
		string(musicFileName), ffmpeg_go.KwArgs{"ac": 2}).Run()

	log.Println(err)

	return musicFileName, nil
}
