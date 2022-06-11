package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/RaymondCode/simple-demo/utils"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/ini.v1"
	"net/http"
	"path"
	"path/filepath"
	"strconv"
)

var (
	url string
)

type VideoListResponse struct {
	entity.Response
	VideoList []entity.Video `json:"video_list"`
}

func init() {
	f, err := ini.Load("./conf.ini")
	if err != nil {
		utils.LogrusObj.Info(err)
		panic("配置文件获取失败")
	}
	url = f.Section("publish").Key("Url").String()
	fmt.Println(url)
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")
	self := utils.ParseToken(token)
	user := service.UserInfo(strconv.Itoa(self.ID), self.ID)
	if user.IsEmpty() {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, entity.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	title := c.PostForm("title")

	//使用本地存储
	ext := path.Ext(filepath.Base(data.Filename))
	f := uuid.NewV4()
	finalName := fmt.Sprintf("%s%s", f, ext)
	fmt.Println(finalName)
	saveFile := filepath.Join("./public/videos/", finalName)

	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, entity.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	name := fmt.Sprintf("%s.jpeg", f)
	fmt.Println(name)

	if flag := utils.ReadFrameAsJpeg(saveFile, 5, name); !flag {
		panic("cover wrong")
	}
	coverUrl := fmt.Sprintf("%sstatic/covers/%s", url, name)
	playUrl := fmt.Sprintf("%sstatic/videos/%s", url, finalName)

	//使用七牛
	//code, playUrl := utils.UploadToQiNiu(data)
	//if code != 0 {
	//	c.JSON(http.StatusOK, entity.Response{
	//		StatusCode: 1,
	//		StatusMsg:  playUrl,
	//	})
	//}
	//coverUrl := fmt.Sprintf("%s?vframe/jpg/offset/0", playUrl)

	video := entity.Video{
		Title:    title,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
		User:     *user,
	}
	service.Publish(&video)

	c.JSON(http.StatusOK, entity.Response{
		StatusCode: 0,
		StatusMsg:  "uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	token := c.Query("token")
	uid := c.Query("user_id")
	t := utils.ParseToken(token)
	self := service.UserInfo(strconv.Itoa(t.ID), t.ID)
	user := service.UserInfo(uid, t.ID)
	if user.IsEmpty() {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	videos := service.VideoList(self.ID, user.ID)
	c.JSON(http.StatusOK, VideoListResponse{
		Response: entity.Response{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		VideoList: videos,
	})
}
