package dao

import (
	"fmt"
	"testing"
)

func TestCreateFollowTx(t *testing.T) {
	CreateFollowTx(2, 1)
}

func TestGetFollow(t *testing.T) {
	fmt.Println(GetFollow(2))
}

func TestGetFollower(t *testing.T) {
	fmt.Println(GetFollower(1))
}
