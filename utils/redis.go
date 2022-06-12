package utils

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/ini.v1"
	"strings"
	"time"
)

var (
	pool        *redis.Pool //创建redis连接池
	maxIdle     int         //最初的连接数量
	maxActive   int         //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
	idleTimeOut string      //连接关闭时间 300秒 （300秒不使用自动关闭）
	url         string      //redis服务器url
	port        string      //端口
)

func init() {
	file, err := ini.Load("./conf.ini")
	if err != nil {
		LogrusObj.Info(err)
		panic("配置文件有误")
	}
	maxIdle, _ = file.Section("redis").Key("maxIdle").Int()
	maxActive, _ = file.Section("redis").Key("maxActive").Int()
	idleTimeOut = file.Section("redis").Key("idleTimeOut").String()
	url = file.Section("redis").Key("url").String()
	port = file.Section("redis").Key("port").String()

	duration, _ := time.ParseDuration(idleTimeOut)
	path := strings.Join([]string{"http://", url, ":", port}, "")
	fmt.Println(path)
	pool = &redis.Pool{ //实例化一个连接池
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: duration,
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			return redis.Dial("tcp", path)
		},
	}
}

func GetConn() redis.Conn {
	return pool.Get()
}
