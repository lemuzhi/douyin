package test

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io"
	"os"
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
