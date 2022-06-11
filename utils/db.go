package utils

import (
	"github.com/RaymondCode/simple-demo/entity"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

var db *gorm.DB

func init() {
	file, err := ini.Load("./conf.ini")
	if err != nil {
		LogrusObj.Info(err)
		panic("配置文件有误")
	}
	DbHost := file.Section("mysql").Key("DbHost").String()
	DbPort := file.Section("mysql").Key("DbPort").String()
	DbUser := file.Section("mysql").Key("DbUser").String()
	DbPassWord := file.Section("mysql").Key("DbPassWord").String()
	DbName := file.Section("mysql").Key("DbName").String()
	path := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=true&loc=Local"}, "")

	dsn := path
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		LogrusObj.Info(err)
		panic("连接数据库失败, error=" + err.Error())
	}

	sqlDB, err := db.DB()

	if err != nil {
		LogrusObj.Info(err)
	}

	// SetMaxIdleCons 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenCons 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(50)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err = db.AutoMigrate(&entity.Video{}, &entity.User{}, &entity.Favorite{}, &entity.Comment{}, &entity.Follow{}); err != nil {
		LogrusObj.Info(err)
	}
}

func GetDB() *gorm.DB {
	return db
}
