package controller

import (
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserLoginResponse struct {
	entity.Response
	UserId uint `json:"user_id,omitempty"`
	Token  uint `json:"token"`
}

type UserResponse struct {
	entity.Response
	User entity.User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	user := entity.User{Name: username, Passwd: password}

	if flag := service.Register(&user); !flag {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: entity.Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: entity.Response{StatusCode: 0},
			UserId:   user.ID,
			Token:    user.ID,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	user := entity.User{Name: username, Passwd: password}
	if flag := service.Login(&user); flag {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: entity.Response{StatusCode: 0, StatusMsg: "success"},
			UserId:   user.ID,
			Token:    user.ID,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: entity.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	uid := c.Query("user_id")
	if user := service.UserInfo(uid, token); !user.IsEmpty() {
		c.JSON(http.StatusOK, UserResponse{
			Response: entity.Response{StatusCode: 0, StatusMsg: "success"},
			User:     *user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: entity.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
