package dao

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/entity"
	"testing"
	"time"
)

func TestCreateVideo(t *testing.T) {
	user := GetUserById(2)
	video := entity.Video{
		Title:    "first edit",
		PlayUrl:  "this is playUrl",
		CoverUrl: "this is coverUrl",
		User:     *user,
	}
	fmt.Println(CreateVideo(&video))
	fmt.Println(video)
}

func TestDeleteVideoById(t *testing.T) {
	video := entity.Video{ID: 2}
	fmt.Println(DeleteVideoById(&video))
	fmt.Println(video)
}

func TestGetVideoByUserId(t *testing.T) {
	videos := GetVideoByUserId(3)
	fmt.Println(videos)
}

func TestGetVideoByTime(t *testing.T) {
	videos := GetVideoByTime(time.Now())
	fmt.Println(videos)
}
