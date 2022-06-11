package service

import (
	"fmt"
	"testing"
)

func TestCommentAction(t *testing.T) {
	CommentCreate(1, "1", "good")
}

func TestCommentDelete(t *testing.T) {
	tx := CommentDelete("7")
	fmt.Println(tx)
}
