package entity

import (
	"gorm.io/gorm"
	"reflect"
	"time"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	ID            uint           `json:"id,omitempty"`
	UserID        uint           `json:"-" gorm:"index:userId"`
	Title         string         `json:"-"`
	PlayUrl       string         `json:"play_url,omitempty" gorm:"type:varchar(100);not null"`
	CoverUrl      string         `json:"cover_url,omitempty" gorm:"type:varchar(100);not null"`
	FavoriteCount uint64         `json:"favorite_count,omitempty"`
	CommentCount  uint64         `json:"comment_count,omitempty"`
	IsFavorite    bool           `json:"is_favorite,omitempty"`
	User          User           `json:"author"`
	CreatedAt     time.Time      `json:"-"`
	UpdatedAt     time.Time      `json:"-"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
}

type Comment struct {
	ID        uint           `json:"id,omitempty"`
	UserID    uint           `json:"-"`
	Content   string         `json:"content,omitempty" gorm:"type:text"`
	VideoId   uint           `json:"-" gorm:"index:videoId"`
	Video     Video          `json:"-"`
	User      User           `json:"user"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type User struct {
	ID            uint           `json:"id,omitempty"`
	Name          string         `json:"name,omitempty" gorm:"type:varchar(20);not null;index:,unique"`
	FollowCount   uint64         `json:"follow_count"`
	FollowerCount uint64         `json:"follower_count"`
	IsFollow      bool           `json:"is_follow"`
	Passwd        string         `json:"-" gorm:"type:varchar(20);not null"`
	CreatedAt     time.Time      `json:"-"`
	UpdatedAt     time.Time      `json:"-"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

type Favorite struct {
	UserID  uint  `json:"user_id,omitempty" gorm:"primary_key;index:userId"`
	VideoID uint  `json:"video_id,omitempty" gorm:"primary_key"`
	User    User  `json:"user"`
	Video   Video `json:"video"`
}

type Follow struct {
	FollowerID uint `json:"follower_id,omitempty" gorm:"primary_key;index:followerId"`
	FollowID   uint `json:"follow_id,omitempty" gorm:"primary_key"`
	Follower   User `json:"follower"`
	Follow     User `json:"follow"`
}

func (f Follow) IsEmpty() bool {
	return reflect.DeepEqual(f, Follow{})
}

func (v Video) IsEmpty() bool {
	return reflect.DeepEqual(v, Video{})
}

func (f Favorite) IsEmpty() bool {
	return reflect.DeepEqual(f, Favorite{})
}

func (u User) IsEmpty() bool {
	return reflect.DeepEqual(u, User{})
}

func (c Comment) IsEmpty() bool {
	return reflect.DeepEqual(c, Comment{})
}
