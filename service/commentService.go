package service

import (
	"github.com/RaymondCode/simple-demo/dao"
	"github.com/RaymondCode/simple-demo/entity"
	"strconv"
)

func CommentCreate(uid uint, vid string, content string) *entity.Comment {
	videoId, _ := strconv.Atoi(vid)
	comment := &entity.Comment{
		Content: content,
		Video: entity.Video{
			ID: uint(videoId),
		},
		User: entity.User{
			ID: uid,
		},
	}
	if user := dao.GetUserById(int(uid)); user.IsEmpty() {
		return &entity.Comment{}
	}
	if video := dao.GetVideoById(videoId); video.IsEmpty() {
		return &entity.Comment{}
	}
	if e := dao.CreateCommentTx(comment); !e {
		return &entity.Comment{}
	}
	return comment
}

func CommentDelete(commentId string) bool {
	id, _ := strconv.Atoi(commentId)
	comment := dao.GetCommentById(id)
	if comment.IsEmpty() {
		return false
	}
	dao.DeleteCommentTx(&comment)
	return true
}

func CommentList(vid string) *[]entity.Comment {
	id, _ := strconv.Atoi(vid)
	return dao.GetCommentByVideoId(id)
}
