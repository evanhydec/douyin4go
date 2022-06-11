package controller

import (
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/RaymondCode/simple-demo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserLoginResponse struct {
	entity.Response
	UserId uint   `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	entity.Response
	User entity.User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	code := utils.GetSHA256HashCode(utils.StringToByteSlice(password))
	user := entity.User{Name: username, Passwd: code}

	if flag := service.Register(&user); !flag {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: entity.Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		token, _ := utils.GenerateToken(int(user.ID), 1)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: entity.Response{StatusCode: 0},
			UserId:   user.ID,
			Token:    token,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	code := utils.GetSHA256HashCode(utils.StringToByteSlice(password))
	user := entity.User{Name: username, Passwd: code}
	if flag := service.Login(&user); flag {
		token, _ := utils.GenerateToken(int(user.ID), 1)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: entity.Response{StatusCode: 0, StatusMsg: "success"},
			UserId:   user.ID,
			Token:    token,
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
	self := utils.ParseToken(token)
	if user := service.UserInfo(uid, self.ID); !user.IsEmpty() {
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
