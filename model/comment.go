package model

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	Id        int64  `json:"id,omitempty" gorm:"primaryKey;AUTO_INCREMENT"`
	UserID    int64  `json:"user_id,omitempty"`
	User      User   `json:"user"`
	VideoID   int64  `json:"video_id"`
	Video     Video  `json:"video"`
	Content   string `json:"content,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
