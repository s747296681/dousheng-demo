package model

import (
	"gorm.io/gorm"
	"time"
)

type Video struct {
	Id            int64  `json:"id,omitempty" gorm:"primaryKey;AUTO_INCREMENT"`
	AuthorID      int64  `json:"author_id"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	VideoName     string `json:"video_name"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	Title         string `json:"title"`
	CommentList   []Comment
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}
