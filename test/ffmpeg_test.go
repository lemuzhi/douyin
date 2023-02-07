package test

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io"
	"os"
	"os/exec"
	"testing"
)

func TestExampleReadFrameAsJpeg(t *testing.T) {
	reader := exampleReadFrameAsJpeg("../public/1fc5884251ca9f20a4fd5f79a458f94f.mp4", 5)
	img, err := imaging.Decode(reader)
	if err != nil {
		t.Fatal(err)
	}
	err = imaging.Save(img, "../public/1.jpeg")
	if err != nil {
		t.Fatal(err)
	}
}

func exampleReadFrameAsJpeg(inFileName string, frameNum int) io.Reader {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		panic(err)
	}
	return buf
}

func TestFfmpegCmd(t *testing.T) {
	cmd := newCmd("ep1.jpg", "../public/1c3c13c2e0c6bd95213377ff07e8dbd2.mp4")

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
}

func newCmd(imageFile, videoFile string) *exec.Cmd {
	return exec.Command("C:\\Users\\muzhi\\Desktop\\douyin\\third_party\\ffmpeg.exe",
		"-i", videoFile, //视频路径
		"-r", "1",
		"-vframes", "1",
		"-q:v", "2",
		"-f", "image2",
		imageFile,
	)
}
