package service

import (
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/entity"
	"strconv"
)

func FollowAction(sid string, mid string, t string) bool {
	follower, _ := strconv.Atoi(sid)
	follow, _ := strconv.Atoi(mid)
	if u1 := dao.GetUserById(follow); u1.IsEmpty() {
		return false
	}
	if u2 := dao.GetUserById(follower); u2.IsEmpty() {
		return false
	}
	if t == "1" {
		if dao.CheckRelation(uint(follower), uint(follow)) {
			return false
		}
		dao.CreateFollowTx(uint(follow), uint(follower))
	} else {
		dao.DeleteFollowTx(uint(follow), uint(follower))
	}
	dao.RedisDeleteUserById(follow)
	dao.RedisDeleteUserById(follower)
	return true
}

func GetFollow(id uint, token string) []entity.User {
	follows := dao.GetFollow(id)
	var res []entity.User

	self, _ := strconv.Atoi(token)
	myFollows := dao.GetFollow(uint(self))
	mfs := make(map[uint]bool, len(myFollows))
	for _, follow := range myFollows {
		mfs[follow.FollowID] = true
	}

	for _, follow := range follows {
		follow.Follow.IsFollow = mfs[follow.FollowID]
		res = append(res, follow.Follow)
	}
	return res
}

func GetFollower(id uint, token string) []entity.User {
	followers := dao.GetFollower(id)
	var res []entity.User

	self, _ := strconv.Atoi(token)
	follows := dao.GetFollow(uint(self))
	fs := make(map[uint]bool, len(follows))
	for _, follow := range follows {
		fs[follow.FollowID] = true
	}

	for _, follow := range followers {
		follow.Follower.IsFollow = fs[follow.FollowerID]
		res = append(res, follow.Follower)
	}
	return res
}
