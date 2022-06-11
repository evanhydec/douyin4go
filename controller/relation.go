package controller

import (
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/service"
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
	user := service.UserInfo(toUser, token)
	if user.IsEmpty() {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}

	t := c.Query("action_type")
	res := service.FollowAction(token, toUser, t)
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
	user := service.UserInfo(uid, token)
	if user.IsEmpty() {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}

	follows := service.GetFollow(user.ID, token)
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
	user := service.UserInfo(uid, token)
	if user.IsEmpty() {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}

	followers := service.GetFollower(user.ID, token)
	c.JSON(http.StatusOK, UserListResponse{
		Response: entity.Response{
			StatusCode: 0,
		},
		UserList: followers,
	})
}
