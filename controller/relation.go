package controller

import (
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/RaymondCode/simple-demo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserListResponse struct {
	entity.Response
	UserList []entity.User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	toUser := c.Query("to_user_id")
	self := utils.ParseToken(token)
	user := service.UserInfo(toUser, self.ID)
	if user.IsEmpty() {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}

	t := c.Query("action_type")
	res := service.FollowAction(self.ID, toUser, t)
	if res {
		c.JSON(http.StatusOK, entity.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		})
	} else {
		c.JSON(http.StatusOK, entity.Response{
			StatusCode: 0,
			StatusMsg:  "you have followed",
		})
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	token := c.Query("token")
	uid := c.Query("user_id")
	self := utils.ParseToken(token)
	user := service.UserInfo(uid, self.ID)
	if user.IsEmpty() {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}

	follows := service.GetFollow(user.ID, self.ID)
	c.JSON(http.StatusOK, UserListResponse{
		Response: entity.Response{
			StatusCode: 0,
		},
		UserList: follows,
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	token := c.Query("token")
	uid := c.Query("user_id")
	self := utils.ParseToken(token)
	user := service.UserInfo(uid, self.ID)
	if user.IsEmpty() {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}

	followers := service.GetFollower(user.ID, self.ID)
	c.JSON(http.StatusOK, UserListResponse{
		Response: entity.Response{
			StatusCode: 0,
		},
		UserList: followers,
	})
}
