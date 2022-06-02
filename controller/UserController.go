package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/RaymondCode/simple-demo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	if username == "" || password == "" {
		c.JSON(http.StatusBadRequest, model.UserLoginResponse{
			Response: model.Response{
				StatusCode: -1,
				StatusMsg:  "some params is lose",
			},
			UserId: -1,
			Token:  "",
		})
		return
	}
	if len(username) > 32 || len(password) > 32 {
		c.JSON(http.StatusBadRequest, model.UserLoginResponse{
			Response: model.Response{
				StatusCode: -1,
				StatusMsg:  "check username or password len",
			},
			UserId: -1,
			Token:  "",
		})
		return
	}
	userId, err := service.CreateUser(username, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.UserLoginResponse{
			Response: model.Response{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
			UserId: -1,
			Token:  "",
		})
		return
	}
	token, err := utils.GenToken(username, password)
	if err != nil {
		fmt.Println("can not generate token")
		c.JSON(http.StatusInternalServerError, model.UserLoginResponse{
			Response: model.Response{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
			UserId: 0,
			Token:  "",
		})
		return
	}
	c.JSON(http.StatusOK, model.UserLoginResponse{
		Response: model.Response{StatusCode: 0},
		UserId:   userId,
		Token:    token,
	})
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	if username == "" || password == "" {
		c.JSON(http.StatusBadRequest, model.UserLoginResponse{
			Response: model.Response{
				StatusCode: -1,
				StatusMsg:  "some params is lose",
			},
			UserId: -1,
			Token:  "",
		})
		return
	}

	userId, token, err := service.ValidateUser(username, password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.UserLoginResponse{
			Response: model.Response{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
			UserId: 0,
			Token:  "",
		})
		return
	}
	c.JSON(http.StatusOK, model.UserLoginResponse{
		Response: model.Response{StatusCode: 0},
		UserId:   userId,
		Token:    token,
	})

}

func GetUserInfo(c *gin.Context) {
	userIdStr := c.Query("user_id")
	token := c.Query("token")
	if userIdStr == "" || token == "" {
		c.JSON(http.StatusBadRequest, model.UserResponse{
			Response: model.Response{
				StatusCode: -1,
				StatusMsg:  "some params is lose",
			},
			User: model.User{},
		})
		return
	}
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.UserResponse{
			Response: model.Response{
				StatusCode: -1,
				StatusMsg:  "please check your userId",
			},
			User: model.User{},
		})
		return
	}
	user, err := service.GetUserInfo(userId, token)
	if err != nil || user == nil {
		c.JSON(http.StatusInternalServerError, model.UserResponse{
			Response: model.Response{
				StatusCode: -1,
				StatusMsg:  err.Error(),
			},
			User: model.User{},
		})
		return
	}
	c.JSON(http.StatusOK, model.UserResponse{
		Response: model.Response{
			StatusCode: 0,
		},
		User: *user,
	})

}
