package dao

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/entity"
	"testing"
)

func TestRedisSetUser(t *testing.T) {
	RedisSetUser(entity.User{
		ID:            2,
		Name:          "范炬涛",
		FollowCount:   5,
		FollowerCount: 3,
		IsFollow:      true,
	})
}

func TestRedisGetUserById(t *testing.T) {
	user, _ := RedisGetUserById(2)
	fmt.Println(user)
}

func TestRedisDeleteUserById(t *testing.T) {
	tx := RedisDeleteUserById(2)
	fmt.Println(tx)
}

//func TestRelation(t *testing.T) {
//	RedisSetRelation(1, 2, true)
//	_, r := RedisGetRelation(1, 2)
//	fmt.Println("关系是：", r)
//	b := RedisDeleteRelation(1, 2)
//	fmt.Println(b)
//}
//
//func TestGetRelation(t *testing.T) {
//	_, relation := RedisGetRelation(1, 2)
//	fmt.Println(relation)
//}
//
//func TestRedisSetPVideos(t *testing.T) {
//	videos := GetVideoByUserId(1)
//	res := RedisSetPVideos(1, videos)
//	fmt.Println(res)
//}
//
//func TestRedisGetPVideosById(t *testing.T) {
//	id, _ := RedisGetPVideosById(1)
//	fmt.Println(id)
//}
//
//func TestRedisSetFVideos(t *testing.T) {
//	videos := GetFavouriteVideoByUserID(1)
//	res := RedisSetFVideos(1, videos)
//	fmt.Println(res)
//}
//
//func TestRedisGetFVideosById(t *testing.T) {
//	id, _ := RedisGetFVideosById(1)
//	fmt.Println(id)
//}
