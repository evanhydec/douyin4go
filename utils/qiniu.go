package utils

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"gopkg.in/ini.v1"
	"mime/multipart"
)

var (
	accessKey string // 秘钥对
	serectKey string
	bucket    string // 空间名称
	imgUrl    string //cdn的url
	savePath  string //空间下视频存放路径
)

func init() {
	file, err := ini.Load("./conf.ini")
	if err != nil {
		LogrusObj.Info(err)
		panic("配置文件有误")
	}
	accessKey = file.Section("qiniu").Key("accessKey").String()
	serectKey = file.Section("qiniu").Key("serectKey").String()
	bucket = file.Section("qiniu").Key("bucket").String()
	imgUrl = file.Section("qiniu").Key("imgUrl").String()
	savePath = file.Section("qiniu").Key("savePath").String()
}

func UploadToQiNiu(file *multipart.FileHeader) (int, string) {

	src, err := file.Open()
	if err != nil {
		return 10011, err.Error()
	}
	defer src.Close()

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, serectKey)

	// 获取上传凭证
	upToken := putPolicy.UploadToken(mac)
	region, _ := storage.GetRegionByID("cn-east-2")
	// 配置参数
	cfg := storage.Config{
		Zone:          &region, // 浙江
		UseCdnDomains: false,
		UseHTTPS:      false, // 非https
	}
	formUploader := storage.NewFormUploader(&cfg)

	ret := storage.PutRet{}        // 上传后返回的结果
	putExtra := storage.PutExtra{} // 额外参数

	// 上传 自定义key，可以指定上传目录及文件名和后缀，
	key := savePath + file.Filename // 上传路径，如果当前目录中已存在相同文件，则返回上传失败错误
	err = formUploader.Put(context.Background(), &ret, upToken, key, src, file.Size, &putExtra)

	// 以默认key方式上传
	// err = formUploader.PutWithoutKey(context.Background(), &ret, upToken, src, fileSize, &putExtra)

	// 自定义key，上传指定路径的文件
	// localFilePath = "./aa.jpg"
	// err = formUploader.PutFile(context.Background(), &ret, upToken, key, localFilePath, &putExtra)

	// 默认key，上传指定路径的文件
	// localFilePath = "./aa.jpg"
	// err = formUploader.PutFile(context.Background(), &ret, upToken, key, localFilePath, &putExtra)

	if err != nil {
		code := 501
		return code, err.Error()
	}

	url := imgUrl + ret.Key // 返回上传后的文件访问路径
	return 0, url
}
