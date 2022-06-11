package service

import (
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/entity"
	"time"
)

func Publish(video *entity.Video) bool {
	return dao.CreateVideo(video)
}

func VideoList(self uint, uid uint) []entity.Video {
	videos := dao.GetVideoByUserId(uid)
	var res []entity.Video

	follows := dao.GetFollow(self)
	fls := make(map[uint]bool, len(follows))
	for _, follow := range follows {
		fls[follow.FollowID] = true
	}

	for _, video := range videos {
		video.User.IsFollow = fls[video.UserID]
		res = append(res, video)
	}

	return res
}

func Feed(token int, timestamp int64) []entity.Video {
	t := time.Unix(timestamp, 0)
	videos := dao.GetVideoByTime(t)
	var res []entity.Video
	if token == 0 {
		return videos
	}

	follows := dao.GetFollow(uint(token))
	follow := make(map[uint]bool, len(follows))
	for _, followe := range follows {
		follow[followe.FollowID] = true
	}

	favourites := dao.GetFavouriteVideoByUserID(uint(token))
	favourite := make(map[uint]bool, len(favourites))
	for _, favourit := range favourites {
		favourite[favourit.ID] = true
	}

	for _, video := range videos {
		video.User.IsFollow = follow[video.UserID]
		video.IsFavorite = favourite[video.ID]
		res = append(res, video)
	}
	return res
}
