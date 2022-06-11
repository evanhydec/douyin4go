package dao

import (
	"encoding/json"
	"fmt"
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/utils"
	"github.com/garyburd/redigo/redis"
)

var (
	slow = 600 //用户信息、关注信息、点赞信息
)

func RedisSetUser(user entity.User) bool {
	conn := utils.GetConn()
	defer conn.Close()

	key := fmt.Sprintf("user:%v", user.ID)

	var data []byte
	var err error

	if data, err = json.Marshal(user); err != nil {
		fmt.Println(err)
		return false
	}

	if _, err = conn.Do("set", key, data, "EX", slow); err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println("redis新增user")
	return true
}

func RedisGetUserById(id int) (*entity.User, bool) {
	conn := utils.GetConn()
	defer conn.Close()

	keys := fmt.Sprintf("user:%v", id)
	var res *entity.User

	r, err := redis.Bytes(conn.Do("get", keys))
	if len(r) > 0 {
		if err = json.Unmarshal(r, &res); err != nil {
			fmt.Println(err)
			return nil, true
		}
		fmt.Println("redis获得user")
		return res, false
	}
	return nil, true
}

func RedisDeleteUserById(id int) bool {
	conn := utils.GetConn()
	defer conn.Close()

	keys := fmt.Sprintf("user:%v", id)
	_, err := conn.Do("del", keys)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println("redis删除user")
	return true
}

//
//func RedisSetRelation(self int, target int, relation bool) bool {
//	conn := utils.GetConn()
//	defer conn.Close()
//	key := fmt.Sprintf("%v:%v", self, target)
//
//	if _, err := conn.Do("set", key, relation, "EX", slow); err != nil {
//		fmt.Println(err)
//		return false
//	}
//	fmt.Println("redis新增关系")
//	return true
//}
//
//func RedisGetRelation(self int, target int) (bool, bool) {
//	conn := utils.GetConn()
//	defer conn.Close()
//
//	key := fmt.Sprintf("%v:%v", self, target)
//	res, err := redis.Bool(conn.Do("get", key))
//	if err != nil {
//		fmt.Println(err)
//		return false, false
//	}
//	fmt.Println("redis获得关系")
//	return true, res
//}
//
//func RedisDeleteRelation(self int, target int) bool {
//	conn := utils.GetConn()
//	defer conn.Close()
//
//	key := fmt.Sprintf("%v:%v", self, target)
//	_, err := conn.Do("del", key)
//	if err != nil {
//		fmt.Println(err)
//		return false
//	}
//	fmt.Println("redis删除关系")
//	return true
//}
//
//// RedisSetPVideos save user's publish videos
//// note videos are mapped by user id rather than video id
//func RedisSetPVideos(id int, videos []entity.Video) bool {
//	conn := utils.GetConn()
//	defer conn.Close()
//
//	key := fmt.Sprintf("pvideos:%v", id)
//	data, err := json.Marshal(videos)
//	_, err = conn.Do("set", key, data, "EX", mid)
//	if err != nil {
//		fmt.Println(err)
//		return false
//	}
//	fmt.Println("redis新增videos")
//	return true
//}
//
//// RedisGetPVideosById get user's publish videos
//// here id is the user's id
//func RedisGetPVideosById(id int) (*[]entity.Video, bool) {
//	conn := utils.GetConn()
//	defer conn.Close()
//
//	key := fmt.Sprintf("pvideos:%v", id)
//	var res *[]entity.Video
//	r, err := redis.Bytes(conn.Do("get", key))
//	if len(r) > 0 {
//		err = json.Unmarshal(r, &res)
//		if err != nil {
//			fmt.Println(err)
//			return nil, false
//		}
//		fmt.Println("redis获得videos")
//		return res, true
//	}
//	return nil, false
//}
//
//// RedisDeletePVideosById delete user's publish videos
//func RedisDeletePVideosById(id int) bool {
//	conn := utils.GetConn()
//	defer conn.Close()
//
//	keys := fmt.Sprintf("pvideos:%v", id)
//	_, err := conn.Do("del", keys)
//	if err != nil {
//		fmt.Println(err)
//		return false
//	}
//	fmt.Println("redis删除videos")
//	return true
//}
//
//// RedisSetFVideos save user's favourite videos
//// note videos are mapped by user id rather than video id
//func RedisSetFVideos(id int, videos []entity.Video) bool {
//	conn := utils.GetConn()
//	defer conn.Close()
//
//	key := fmt.Sprintf("fvideos:%v", id)
//	data, err := json.Marshal(videos)
//	_, err = conn.Do("set", key, data, "EX", mid)
//	if err != nil {
//		fmt.Println(err)
//		return false
//	}
//	fmt.Println("redis新增fvideos")
//	return true
//}
//
//// RedisGetFVideosById get user's favourite videos
//// here id is the user's id
//func RedisGetFVideosById(id int) (*[]entity.Video, bool) {
//	conn := utils.GetConn()
//	defer conn.Close()
//
//	key := fmt.Sprintf("fvideos:%v", id)
//	var res *[]entity.Video
//	r, err := redis.Bytes(conn.Do("get", key))
//	if len(r) > 0 {
//		err = json.Unmarshal(r, &res)
//		if err != nil {
//			fmt.Println(err)
//			return nil, false
//		}
//		fmt.Println("redis获得fvideos")
//		return res, true
//	}
//	return nil, false
//}
//
//// RedisDeleteFVideosById delete user's favourite videos
//func RedisDeleteFVideosById(id int) bool {
//	conn := utils.GetConn()
//	defer conn.Close()
//
//	keys := fmt.Sprintf("fvideos:%v", id)
//	_, err := conn.Do("del", keys)
//	if err != nil {
//		fmt.Println(err)
//		return false
//	}
//	fmt.Println("redis删除fvideos")
//	return true
//}
//
//func RedisSetComments(id int, comments []entity.Comment) bool {
//	conn := utils.GetConn()
//	defer conn.Close()
//
//	key := fmt.Sprintf("comments:%v", id)
//	data, err := json.Marshal(comments)
//	_, err = conn.Do("set", key, data, "EX", mid)
//	if err != nil {
//		fmt.Println(err)
//		return false
//	}
//	fmt.Println("redis新增comments")
//	return true
//}
//
//func RedisGetCommentsById(id int) (*[]entity.Comment, bool) {
//	conn := utils.GetConn()
//	defer conn.Close()
//
//	key := fmt.Sprintf("comments:%v", id)
//	var res *[]entity.Comment
//	r, err := redis.Bytes(conn.Do("get", key))
//	if len(r) > 0 {
//		err = json.Unmarshal(r, &res)
//		if err != nil {
//			fmt.Println(err)
//			return nil, false
//		}
//		fmt.Println("redis获得comments")
//		return res, true
//	}
//	return nil, false
//}
//
//func RedisDeleteCommentsById(id int) bool {
//	conn := utils.GetConn()
//	defer conn.Close()
//
//	keys := fmt.Sprintf("comments:%v", id)
//	_, err := conn.Do("del", keys)
//	if err != nil {
//		fmt.Println(err)
//		return false
//	}
//	fmt.Println("redis删除comments")
//	return true
//}
//
//func RedisSetFollows(id int, follows []entity.Follow) bool {
//	conn := utils.GetConn()
//	defer conn.Close()
//
//	key := fmt.Sprintf("follows:%v", id)
//	data, err := json.Marshal(follows)
//	_, err = conn.Do("set", key, data, "EX", slow)
//	if err != nil {
//		fmt.Println(err)
//		return false
//	}
//	fmt.Println("redis新增follows")
//	return true
//}
//
//func RedisGetFollowsById(id int) (*[]entity.Follow, bool) {
//	conn := utils.GetConn()
//	defer conn.Close()
//
//	key := fmt.Sprintf("follows:%v", id)
//	var res *[]entity.Follow
//	r, err := redis.Bytes(conn.Do("get", key))
//	if len(r) > 0 {
//		err = json.Unmarshal(r, &res)
//		if err != nil {
//			fmt.Println(err)
//			return nil, false
//		}
//		fmt.Println("redis获得follows")
//		return res, true
//	}
//	return nil, false
//}
//
//func RedisDeleteFollowsById(id int) bool {
//	conn := utils.GetConn()
//	defer conn.Close()
//
//	keys := fmt.Sprintf("follows:%v", id)
//	_, err := conn.Do("del", keys)
//	if err != nil {
//		fmt.Println(err)
//		return false
//	}
//	fmt.Println("redis删除follows")
//	return true
//}
