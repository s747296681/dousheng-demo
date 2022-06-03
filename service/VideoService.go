package service

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/repository"
	"github.com/RaymondCode/simple-demo/utils"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"time"
)

//后续优化相关逻辑
func UploadVideo(ctx *gin.Context, data *multipart.FileHeader) error {
	fileKey := utils.GenerateVideoKey(data.Filename)
	open, err := data.Open()
	if err != nil {
		return err
	}
	err = utils.Upload(fileKey, open)
	if err != nil {
		return err
	}
	claim, err := utils.GetClaimInfoByCtx(ctx)
	if err != nil {
		return err
	}
	user, err := repository.GetUserByUserName(claim.UserName)
	if err != nil {
		return err
	}
	createdVideo := &model.Video{
		AuthorID:      user.Id,
		PlayUrl:       utils.GetVideoUrl(fileKey),
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		VideoName:     fileKey,
		IsFavorite:    false,
		CommentList:   nil,
	}
	_, err = repository.CreateVideo(createdVideo)
	if err != nil {
		return err
	}
	return nil
}

func GetVideos(queryTime time.Time) ([]model.Video, int64, error) {
	videos, err := repository.GetVideosByTime(queryTime)
	if err != nil {
		return nil, -1, err
	}
	if len(videos) == 0 {
		return videos, time.Now().UnixMilli(), nil
	}
	lastVideo := videos[len(videos)-1]
	return videos, lastVideo.CreatedAt.Unix(), nil
}

func GetVideosByUserId(id int64) ([]model.Video, error) {
	return repository.GetVideosByUserId(id)
}
