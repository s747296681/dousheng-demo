package controller

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/RaymondCode/simple-demo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func Publish(c *gin.Context) {

	//if _, exist := usersLoginInfo[token]; !exist {
	//	c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//	return
	//}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	err = service.UploadVideo(c, data)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	//filename := filepath.Base(data.Filename)
	//user := usersLoginInfo[token]
	//finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	//saveFile := filepath.Join("./public/", "finalName")
	//if err := c.SaveUploadedFile(data, saveFile); err != nil {
	//	c.JSON(http.StatusOK, model.Response{
	//		StatusCode: 1,
	//		StatusMsg:  err.Error(),
	//	})
	//	return
	//}

	c.JSON(http.StatusOK, model.Response{
		StatusCode: 0,
		StatusMsg:  "video uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	idStr := c.Query("user_id")
	if idStr == "" {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "some params is missing",
		})
		return
	}
	userId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	videos, err := service.GetVideosByUserId(userId)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model.VideoListResponse{
		Response: model.Response{
			StatusCode: 0,
		},
		VideoList: videos,
	})
}

func Feed(c *gin.Context) {
	paramTime := c.Query("latest_time")
	var queryTime time.Time
	if paramTime == "" {
		queryTime = time.Now()
	} else {
		ctime, err := utils.ParseStringTime(paramTime)
		if err != nil {
			c.JSON(http.StatusOK, model.FeedResponse{
				Response:  model.Response{StatusCode: -1},
				VideoList: nil,
				NextTime:  time.Now().Unix(),
			})
		}
		queryTime = ctime
	}
	videos, times, err := service.GetVideos(queryTime)
	if err != nil {
		c.JSON(http.StatusOK, model.FeedResponse{
			Response:  model.Response{StatusCode: -1},
			VideoList: nil,
			NextTime:  time.Now().Unix(),
		})
	}
	c.JSON(http.StatusOK, model.FeedResponse{
		Response:  model.Response{StatusCode: 0},
		VideoList: videos,
		NextTime:  times,
	})
}
