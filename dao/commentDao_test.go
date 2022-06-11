package dao

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/entity"
	"testing"
)

func TestGetCommentByVideoId(t *testing.T) {
	comment := GetCommentByVideoId(4)
	fmt.Println(comment)
}

func TestDeleteComment(t *testing.T) {
	DeleteComment(&entity.Comment{ID: 1})
}
