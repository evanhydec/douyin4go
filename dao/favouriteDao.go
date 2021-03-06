package dao

import (
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/utils"
	"gorm.io/gorm"
)

func VideoFavouriteTx(t bool, id int, vid int) bool {
	tx := utils.GetDB().Begin()
	if t {
		if err := tx.Model(entity.Video{}).Set("gorm:query_option", "FOR UPDATE").Where("id = ? ", vid).Update("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
			utils.LogrusObj.Info(err)
			tx.Rollback()
			return false
		}
		if e := tx.Set("gorm:query_option", "FOR UPDATE").Create(&entity.Favorite{UserID: uint(id), VideoID: uint(vid)}).Error; e != nil {
			utils.LogrusObj.Info(e)
			tx.Rollback()
			return false
		}
	} else {
		if err := tx.Model(entity.Video{}).Where("id = ? ", vid).Set("gorm:query_option", "FOR UPDATE").Update("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error; err != nil {
			utils.LogrusObj.Info(err)
			tx.Rollback()
			return false
		}
		if err := tx.Where("user_id = ? and video_id = ?", id, vid).Set("gorm:query_option", "FOR UPDATE").Delete(entity.Favorite{}).Error; err != nil {
			utils.LogrusObj.Info(err)
			tx.Rollback()
			return false
		}
	}
	tx.Commit()
	return true
}

func GetFavourite(vid uint, uid uint) bool {
	db := utils.GetDB()
	var res entity.Favorite
	db.Where("user_id = ? and video_id = ?", uid, vid).Find(&res)
	return !res.IsEmpty()
}

func GetFavouriteVideoByUserID(uid uint) []entity.Video {
	if uid == 0 {
		return nil
	}
	db := utils.GetDB()
	var favs []entity.Favorite
	db.Preload("Video").Where("user_id = ?", uid).Find(&favs)

	var res []entity.Video
	for _, fav := range favs {
		res = append(res, fav.Video)
	}
	return res
}
