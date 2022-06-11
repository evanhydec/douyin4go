package service

import (
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/entity"
	"strconv"
)

func Register(user *entity.User) bool {
	if dao.GetUserByName(user.Name) {
		return false
	}
	dao.CreateUser(user)
	return true
}

func Login(user *entity.User) bool {
	return dao.GetUserByNameAndPwd(user)
}

func UserInfo(uid string, token int) *entity.User {
	if uid == "0" {
		return &entity.User{}
	}

	user, _ := strconv.Atoi(uid)

	res := dao.GetUserById(user)
	res.IsFollow = dao.CheckRelation(uint(token), uint(user))

	return res
}
