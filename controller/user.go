package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync/atomic"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]model.User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

var userIdSequence = int64(1)

type UserLoginResponse struct {
	utils.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	utils.Response
	User model.User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token, err := utils.GenToken(username, password)
	if err != nil {
		fmt.Println("can not generate token")
		c.JSON(http.StatusInternalServerError, UserLoginResponse{
			Response: utils.Response{},
			UserId:   0,
			Token:    "",
		})
	}

	if _, exist := usersLoginInfo[username]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: utils.Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		atomic.AddInt64(&userIdSequence, 1)
		newUser := model.User{
			Id:   userIdSequence,
			Name: username,
		}
		usersLoginInfo[username] = newUser
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: utils.Response{StatusCode: 0},
			UserId:   userIdSequence,
			Token:    token,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	if user, exist := usersLoginInfo[username]; exist {
		token, err := utils.GenToken(username, password)
		if err != nil {
			fmt.Println("can not generate token")
			c.JSON(http.StatusInternalServerError, UserLoginResponse{
				Response: utils.Response{},
				UserId:   0,
				Token:    "",
			})
		}

		c.JSON(http.StatusOK, UserLoginResponse{
			Response: utils.Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: utils.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	ctx, err := utils.GetClaimInfoByCtx(c)
	if err != nil {
		return
	}

	if user, exist := usersLoginInfo[ctx.UserName]; exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: utils.Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: utils.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
