package repository

import (
	"github.com/RaymondCode/simple-demo/model"
	"time"
)

func CreateVideo(video *model.Video) (int64, error) {
	return video.Id, DB.Create(video).Error
}

func GetVideosByTime(lastTime time.Time) ([]model.Video, error) {
	videos := []model.Video{}
	return videos, DB.Preload("Author").Order("created_at desc").Where("created_at < ?", &lastTime).Limit(30).Find(&videos).Error
}

func GetVideosByUserId(id int64) ([]model.Video, error) {
	videos := []model.Video{}
	return videos, DB.Preload("Author").Order("created_at desc").Where("author_id = ?", &id).Find(&videos).Error
}
