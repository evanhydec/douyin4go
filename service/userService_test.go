package service

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/entity"
	"strconv"
	"testing"
	"time"
)

func TestLogin(t *testing.T) {
}

func TestRegister(t *testing.T) {
	for i := 0; i < 100; i++ {
		go func(n int) {
			user := entity.User{Name: strconv.Itoa(n), Passwd: "123123"}
			Register(&user)
		}(i)
	}
	time.Sleep(time.Second * 5)
}

func TestUserInfo(t *testing.T) {
	fmt.Println(UserInfo("5", 5))
}
