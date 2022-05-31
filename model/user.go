package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id            int64  `json:"id,omitempty" gorm:"primaryKey;AUTO_INCREMENT"`
	UserName      string `json:"user_name,omitempty"`
	Name          string `json:"name,omitempty"`
	PassWord      string `json:"pass_word,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
	CommentList   []Comment
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}
