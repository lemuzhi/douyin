package utils

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"log"
	"os"
)

func FfmpegCoverJpeg(playUrl, coverUrl string, frameNum int) {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(playUrl).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		panic(err)
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		log.Println(err)
	}
	err = imaging.Save(img, coverUrl)
	if err != nil {
		log.Println(err)
	}
}
