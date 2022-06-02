package controller

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/gin-gonic/gin"
)

//type User struct {
//	UserDataService service.UserService
//}

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

//func Login(c *gin.Context) {
//	username := c.Query("username")
//	password := c.Query("password")
//
//	if user, exist := usersLoginInfo[username]; exist {
//		token, err := utils.GenToken(username, password)
//		if err != nil {
//			fmt.Println("can not generate token")
//			c.JSON(http.StatusInternalServerError, model.UserLoginResponse{
//				Response: model.Response{},
//				UserId:   0,
//				Token:    "",
//			})
//		}
//
//		c.JSON(http.StatusOK, model.UserLoginResponse{
//			Response: model.Response{StatusCode: 0},
//			UserId:   user.Id,
//			Token:    token,
//		})
//	} else {
//		c.JSON(http.StatusOK, model.UserLoginResponse{
//			Response: model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
//		})
//	}
//}

func Test(c *gin.Context) {

	//dataService := service.NewUserDataService(repository.NewUserRepository(repository.GetDB()))
	//user, err := service.CreateUser(model.User{})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(user)
	//
	//c.JSON(http.StatusOK, model.UserLoginResponse{
	//	Response: model.Response{StatusCode: 0},
	//	UserId:   1,
	//	Token:    "asasasas",
	//})
}

//func UserInfo(c *gin.Context) {
//	ctx, err := utils.GetClaimInfoByCtx(c)
//	if err != nil {
//		return
//	}
//
//	if user, exist := usersLoginInfo[ctx.UserName]; exist {
//		c.JSON(http.StatusOK, model.UserResponse{
//			Response: model.Response{StatusCode: 0},
//			User:     user,
//		})
//	} else {
//		c.JSON(http.StatusOK, model.UserResponse{
//			Response: model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
//		})
//	}
//}
