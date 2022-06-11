package service

import (
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/entity"
	"strconv"
)

func FavouriteAction(videoId string, uid uint, t string) bool {
	vid, _ := strconv.Atoi(videoId)
	if user := dao.GetUserById(int(uid)); user.IsEmpty() {
		return false
	}
	if video := dao.GetVideoById(vid); video.IsEmpty() {
		return false
	}
	if t == "1" && dao.GetFavourite(uint(vid), uid) {
		return false
	}
	return dao.VideoFavouriteTx(t == "1", int(uid), vid)
}

func FavouriteList(uid uint, token int) []entity.Video {

	follows := dao.GetFollow(uint(token))
	follow := make(map[uint]bool, len(follows))
	for _, followe := range follows {
		follow[followe.FollowID] = true
	}

	var videos []entity.Video
	favourites := dao.GetFavouriteVideoByUserID(uid)
	for _, favourite := range favourites {
		favourite.User.IsFollow = follow[favourite.UserID]
		videos = append(videos, favourite)
	}
	return videos
}
