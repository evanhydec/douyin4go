package dao

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/utils"
	"gorm.io/gorm"
)

func CreateFollowTx(mid uint, sid uint) bool {
	tx := utils.GetDB().Begin()
	if err := tx.Model(entity.User{}).Where("id = ? ", sid).Update("follow_count", gorm.Expr("follow_count + ?", 1)).Error; err != nil {
		fmt.Println("follower修改出错，进行回滚")
		tx.Rollback()
		return false
	}
	if err := tx.Model(entity.User{}).Where("id = ? ", mid).Update("follower_count", gorm.Expr("follower_count + ?", 1)).Error; err != nil {
		fmt.Println("follow修改出错，进行回滚")
		tx.Rollback()
		return false
	}
	if err := tx.Create(&entity.Follow{FollowerID: sid, FollowID: mid}).Error; err != nil {
		fmt.Println("创建follow记录出错，进行回滚")
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

func DeleteFollowTx(mid uint, sid uint) bool {
	tx := utils.GetDB().Begin()
	if err := tx.Where("follower_id = ? and follow_id = ?", sid, mid).Delete(entity.Follow{}).Error; err != nil {
		fmt.Println("删除follow记录出错，进行回滚")
		tx.Rollback()
		return false
	}
	if err := tx.Model(entity.User{}).Where("id = ? ", sid).Update("follow_count", gorm.Expr("follow_count - ?", 1)).Error; err != nil {
		fmt.Println("follower修改出错，进行回滚")
		tx.Rollback()
		return false
	}
	if err := tx.Model(entity.User{}).Where("id = ? ", mid).Update("follower_count", gorm.Expr("follower_count - ?", 1)).Error; err != nil {
		fmt.Println("follow修改出错，进行回滚")
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

func CheckRelation(sid uint, mid uint) bool {
	if sid == 0 {
		return false
	}
	db := utils.GetDB()
	var res entity.Follow
	db.Where("follower_id = ? and follow_id = ?", sid, mid).Find(&res)
	return !res.IsEmpty()
}

func GetFollow(sid uint) []entity.Follow {
	if sid == 0 {
		return nil
	}
	db := utils.GetDB()
	var res []entity.Follow
	db.Preload("Follow").Where("follower_id = ?", sid).Find(&res)
	return res
}

func GetFollower(mid uint) []entity.Follow {
	if mid == 0 {
		return nil
	}
	db := utils.GetDB()
	var res []entity.Follow
	db.Preload("Follower").Where("follow_id = ?", mid).Find(&res)
	return res
}
