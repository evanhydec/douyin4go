package dao

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/entity"
	"testing"
)

func TestGetUserById(t *testing.T) {
	fmt.Println(GetUserById(3))
}

func TestGetUserByNameAndPwd(t *testing.T) {
	user := entity.User{Name: "zhangsan", Passwd: "123456"}
	fmt.Println(GetUserByNameAndPwd(&user))
	fmt.Println(user)
}

func TestCreateUser(t *testing.T) {
	user := entity.User{Name: "zhangsan", Passwd: "123123"}
	CreateUser(&user)
	fmt.Println(user)
}

func TestDeleteUser(t *testing.T) {
	fmt.Println(DeleteUser(&entity.User{ID: 1}))
}
