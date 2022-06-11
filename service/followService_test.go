package service

import (
	"fmt"
	"testing"
)

func TestFollowAction(t *testing.T) {
	FollowAction("1", "2", "2")
}

func TestGetFollow(t *testing.T) {
	fmt.Println(GetFollow(1, "1"))
}

func TestGetFollower(t *testing.T) {
	fmt.Println(GetFollower(2, "2"))
}
