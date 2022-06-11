package dao

import (
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/utils"
)

func GetUserById(id int) *entity.User {
	if id == 0 {
		return nil
	}
	if user, b := RedisGetUserById(id); !b {
		return user
	}
	var user entity.User
	db := utils.GetDB()
	db.Where("id = ?", id).First(&user)
	RedisSetUser(user)
	return &user
}

func GetUserByNameAndPwd(user *entity.User) bool {
	db := utils.GetDB()
	tx := db.Where("name = ? and passwd = ?", user.Name, user.Passwd).First(&user)
	if tx.Error != nil {
		return false
	}
	return true
}

func GetUserByName(name string) bool {
	db := utils.GetDB()
	res := entity.User{}
	db.Where("name = ?", name).Find(&res)
	return !res.IsEmpty()
}

func CreateUser(user *entity.User) {
	db := utils.GetDB()
	db.Create(&user)
}

func DeleteUser(user *entity.User) bool {
	db := utils.GetDB()
	result := db.Delete(user)
	affected := result.RowsAffected
	if affected == 0 {
		return false
	}
	return true
}

// Deprecated: the func is deprecated
func UpdateUser(user *entity.User) bool {
	db := utils.GetDB()
	save := db.Save(&user)
	if save.Error != nil {
		return false
	}
	return true
}
