package service

import (
	"fmt"
	"testing"
	"time"
)

func TestFavouriteAction(t *testing.T) {
	for i := 0; i < 100; i++ {
		go func(id int) {
			FavouriteAction("1", uint(id), "0")
		}(i + 1)
	}
	time.Sleep(time.Second * 5)
}

func TestFavouriteActionTx(t *testing.T) {
	FavouriteAction("1", 2, "1")
}

func TestFavouriteList(t *testing.T) {
	fmt.Println(FavouriteList(0, "2"))
}
