package service

import (
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/entity"
	"strconv"
)

func Register(user *entity.User) bool {
	if dao.GetUserByNameAndPwd(user) {
		return false
	}
	dao.CreateUser(user)
	return true
}

func Login(user *entity.User) bool {
	return dao.GetUserByNameAndPwd(user)
}

func UserInfo(uid string, token string) *entity.User {
	if uid == "" {
		return &entity.User{}
	}

	user, _ := strconv.Atoi(uid)
	self, _ := strconv.Atoi(token)

	res := dao.GetUserById(user)
	res.IsFollow = dao.CheckRelation(uint(self), uint(user))

	return res
}
