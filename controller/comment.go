package controller

import (
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/RaymondCode/simple-demo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CommentListResponse struct {
	entity.Response
	CommentList []entity.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	entity.Response
	Comment entity.Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	self := utils.ParseToken(token)
	user := service.UserInfo(token, self.ID)
	if user.IsEmpty() {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	actionType := c.Query("action_type")
	if actionType == "1" {
		text := c.Query("comment_text")
		vid := c.Query("video_id")
		comment := service.CommentCreate(user.ID, vid, text)
		comment.User = *user
		c.JSON(http.StatusOK,
			CommentActionResponse{
				Response: entity.Response{
					StatusCode: 0,
					StatusMsg:  "success"},
				Comment: *comment,
			})
	} else {
		cid := c.Query("comment_id")
		flag := service.CommentDelete(cid)
		if flag {
			c.JSON(http.StatusOK, entity.Response{StatusCode: 0})
		}
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	id := c.Query("video_id")
	comments := service.CommentList(id)
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    entity.Response{StatusCode: 0},
		CommentList: *comments,
	})
}
