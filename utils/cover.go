package utils

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"os"
)

func ReadFrameAsJpeg(inFileName string, frameNum int, filename string) bool {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		LogrusObj.Info(err)
		return false
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		LogrusObj.Info(err)
		return false
	}

	err = imaging.Save(img, fmt.Sprintf("./public/covers/%s", filename))
	if err != nil {
		LogrusObj.Info(err)
		return false
	}
	return true
}
