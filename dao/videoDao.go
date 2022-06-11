package dao

import (
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/utils"
	"time"
)

func CreateVideo(video *entity.Video) bool {
	db := utils.GetDB()
	create := db.Create(video)
	if create.Error != nil {
		return false
	}
	return true
}

func GetVideoById(vid int) entity.Video {
	db := utils.GetDB()
	var res entity.Video
	db.Where("id = ?", vid).First(&res)
	return res
}

func DeleteVideoById(video *entity.Video) bool {
	db := utils.GetDB()
	tx := db.Where("id = ?", video.ID).Delete(video)
	if tx.Error != nil {
		return false
	}
	return true
}

// Deprecated: it has been deprecated,update should use other func to optimize
func UpdateVideo(video *entity.Video) bool {
	db := utils.GetDB()
	save := db.Save(&video)
	if save.Error != nil {
		return false
	}
	return true
}

func GetVideoByUserId(id uint) []entity.Video {
	if id == 0 {
		return nil
	}
	db := utils.GetDB()
	var videos []entity.Video
	db.Preload("User").Where("user_id = ?", id).Find(&videos)
	return videos
}

func GetVideoByTime(time time.Time) []entity.Video {
	db := utils.GetDB()
	var videos []entity.Video
	db.Preload("User").Where("created_at < ?", time).Order("created_at desc").Limit(30).Find(&videos)
	return videos
}
