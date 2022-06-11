package controller

import (
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/RaymondCode/simple-demo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type FeedResponse struct {
	entity.Response
	VideoList []entity.Video `json:"video_list,omitempty"`
	NextTime  int64          `json:"next_time,omitempty"`
}

//Feed same demo video list for every request
func Feed(c *gin.Context) {
	latest := c.Query("latest_time")
	timestamp := time.Now().Unix()
	if len(latest) == 10 {
		var err error
		timestamp, err = strconv.ParseInt(latest, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, entity.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			})
		}
	}

	token := c.Query("token")

	self := utils.ParseToken(token)

	videos := service.Feed(self.ID, timestamp)

	c.JSON(http.StatusOK, FeedResponse{
		Response:  entity.Response{StatusCode: 0},
		VideoList: videos,
		NextTime:  time.Now().Unix(),
	})
}
