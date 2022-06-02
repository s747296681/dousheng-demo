package service

import (
	"errors"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/repository"
	"github.com/RaymondCode/simple-demo/utils"
	"gorm.io/gorm"
	"time"
)

//type UserService interface {
//	CreateUser(user model.User) (int64, error)
//}
//
//type IUserService struct {
//	UserRepository repository.UserRepository
//}

func CreateUser(username string, password string) (int64, error) {
	pwdByte, err := utils.GeneratePassword(password)
	if err != nil {
		return -1, err
	}
	createdUser := &model.User{
		Id:            0,
		UserName:      "",
		Name:          "HandSomeBoy",
		PassWord:      "",
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
		CommentList:   nil,
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     gorm.DeletedAt{},
	}
	createdUser.UserName = username
	createdUser.PassWord = string(pwdByte)

	return repository.CreateUser(createdUser)
}

func ValidateUser(username string, password string) (int64, string, error) {
	user, err := repository.GetUserByUserName(username)
	if err != nil {
		return -1, "", err
	}
	_, err = utils.ValidatePassword(password, user.PassWord)
	if err != nil {
		return -1, "", err
	}

	token, err := utils.GenToken(username, password)
	if err != nil {
		return -1, "", err
	}

	return user.Id, token, nil
}

func GetUserInfo(userId int64, token string) (*model.User, error) {
	claims, err := utils.ParseToken(token)
	if err != nil {
		return nil, err
	}
	userName := claims.UserName
	user, err := repository.GetUserByUserName(userName)
	if err != nil {
		return nil, err
	}
	if user.Id != userId {
		return nil, errors.New("invalid token")
	}
	return user, nil
}

//func NewUserDataService(userRepository repository.UserRepository) UserService {
//	return IUserService{UserRepository: userRepository}
//}
