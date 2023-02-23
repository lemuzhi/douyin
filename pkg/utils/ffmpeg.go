package utils

import (
	"log"
	"os/exec"
	"runtime"
)

func RunFfmpegCoverJpeg(videoUrl, coverUrl string) {

	var ffmpegPath string

	//获取操作系统类型
	sysType := runtime.GOOS
	switch sysType {
	case "linux":
		ffmpegPath = "./third_party/ffmpeg"
	case "windows":
		ffmpegPath = "./third_party/ffmpeg.exe"
	default:
		ffmpegPath = "./third_party/ffmpeg"
	}

	cmd := exec.Command(ffmpegPath,
		"-i", videoUrl, //视频路径
		"-r", "1", //每秒提取的帧数
		"-vframes", "1", //抽取帧数，抽取1张
		"-q:v", "2", //图片质量
		"-f", "image2", //图片格式
		coverUrl, //图片路径
	)
	if err := cmd.Run(); err != nil {
		log.Println(err)
	}
}
