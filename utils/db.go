package utils

import (
	"github.com/RaymondCode/simple-demo/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

func init() {
	dsn := "root:123123@tcp(localhost:3306)/douyinlocal?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	sqlDB, err := db.DB()

	// SetMaxIdleCons 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenCons 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(50)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	db.AutoMigrate(&entity.Video{}, &entity.User{}, &entity.Favorite{}, &entity.Comment{}, &entity.Follow{})
}

func GetDB() *gorm.DB {
	return db
}
