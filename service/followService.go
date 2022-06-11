package service

import (
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/entity"
	"strconv"
)

func FollowAction(sid int, mid string, t string) bool {
	follow, _ := strconv.Atoi(mid)
	if u1 := dao.GetUserById(follow); u1.IsEmpty() {
		return false
	}
	if u2 := dao.GetUserById(sid); u2.IsEmpty() {
		return false
	}
	if t == "1" {
		if dao.CheckRelation(uint(sid), uint(follow)) {
			return false
		}
		dao.CreateFollowTx(uint(follow), uint(sid))
	} else {
		dao.DeleteFollowTx(uint(follow), uint(sid))
	}
	dao.RedisDeleteUserById(follow)
	dao.RedisDeleteUserById(sid)
	return true
}

func GetFollow(id uint, token int) []entity.User {
	follows := dao.GetFollow(id)
	var res []entity.User

	myFollows := dao.GetFollow(uint(token))
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

func GetFollower(id uint, token int) []entity.User {
	followers := dao.GetFollower(id)
	var res []entity.User

	follows := dao.GetFollow(uint(token))
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
