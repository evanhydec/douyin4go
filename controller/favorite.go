package controller

import (
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/RaymondCode/simple-demo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	self := utils.ParseToken(token)
	user := service.UserInfo(strconv.Itoa(self.ID), self.ID)
	if user.IsEmpty() {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	id := c.Query("video_id")
	t := c.Query("action_type")
	res := service.FavouriteAction(id, user.ID, t)

	if !res {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 0, StatusMsg: "you have liked the video"})
		return
	} else {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 0, StatusMsg: "success"})
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	self := utils.ParseToken(token)
	uid := c.Query("user_id")
	user := service.UserInfo(uid, self.ID)
	if user.IsEmpty() {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}
	videos := service.FavouriteList(user.ID, token)
	c.JSON(http.StatusOK, VideoListResponse{
		Response: entity.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		VideoList: videos,
	})
}
