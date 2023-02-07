package utils

import (
	"log"
	"os/exec"
)

//func FfmpegCoverJpeg(playUrl, coverUrl string, frameNum int) {
//	buf := bytes.NewBuffer(nil)
//	err := ffmpeg.Input(playUrl).
//		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
//		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
//		WithOutput(buf, os.Stdout).
//		Run()
//	if err != nil {
//		panic(err)
//	}
//
//	img, err := imaging.Decode(buf)
//	if err != nil {
//		log.Println(err)
//	}
//	err = imaging.Save(img, coverUrl)
//	if err != nil {
//		log.Println(err)
//	}
//}

func RunFfmpegCoverJpeg(videoUrl, coverUrl string) {
	cmd := exec.Command("./third_party/ffmpeg.exe",
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
