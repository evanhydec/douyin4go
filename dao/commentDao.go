package dao

import (
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/utils"
	"gorm.io/gorm"
)

func DeleteComment(comment *entity.Comment) int {
	db := utils.GetDB()
	var com entity.Comment
	db.Where("id = ?", comment.ID).First(&com)
	if err := db.Delete(&comment).Error; err != nil {
		utils.LogrusObj.Info(err)
		return 0
	}
	VideoComment(false, int(comment.Video.ID))
	return int(com.VideoId)
}

// VideoComment used when user comment the video
func VideoComment(t bool, id int) bool {
	db := utils.GetDB()
	if t {
		if err := db.Model(entity.Video{}).Where("id = ? ", id).Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error; err != nil {
			utils.LogrusObj.Info(err)
			return false
		}
	} else {
		if err := db.Model(entity.Video{}).Where("id = ? ", id).Update("comment_count", gorm.Expr("comment_count - ?", 1)).Error; err != nil {
			utils.LogrusObj.Info(err)
			return false
		}
	}
	return true
}

func CreateCommentTx(comment *entity.Comment) bool {
	tx := utils.GetDB().Begin()
	if err := tx.Create(&comment).Error; err != nil {
		utils.LogrusObj.Info(err)
		tx.Rollback()
		return false
	}
	if err := tx.Model(entity.Video{}).Where("id = ? ", comment.Video.ID).Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error; err != nil {
		utils.LogrusObj.Info(err)
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

func DeleteCommentTx(comment *entity.Comment) int {
	tx := utils.GetDB().Begin()
	id := comment.VideoId
	if err := tx.Delete(&comment).Error; err != nil {
		utils.LogrusObj.Info(err)
		tx.Rollback()
		return 0
	}
	if err := tx.Model(entity.Video{}).Where("id = ? ", id).Update("comment_count", gorm.Expr("comment_count - ?", 1)).Error; err != nil {
		utils.LogrusObj.Info(err)
		tx.Rollback()
		return 0
	}
	tx.Commit()
	return int(id)
}

func GetCommentByVideoId(id int) *[]entity.Comment {
	db := utils.GetDB()
	var comments *[]entity.Comment
	db.Where("video_id = ?", id).Order("created_at desc").Preload("User").Find(&comments)
	return comments
}

func GetCommentById(id int) entity.Comment {
	db := utils.GetDB()
	var res entity.Comment
	db.Where("id = ?", id).Find(&res)
	return res
}
