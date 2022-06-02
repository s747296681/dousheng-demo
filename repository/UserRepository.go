package repository

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/model"
)

//
//type UserRepository interface {
//	CreateUser(user model.User) (int64, error)
//}
//
//type IUserRepository struct {
//	mysqlDB *gorm.DB
//}
//
//func NewUserRepository(db *gorm.DB) UserRepository {
//	return IUserRepository{mysqlDB: db}
//}

func CreateUser(user *model.User) (int64, error) {
	return user.Id, DB.Create(user).Error
}

func GetUserByUserName(userName string) (user *model.User, err error) {
	findingUser := &model.User{}
	result := DB.Where("user_name=?", userName).First(findingUser)
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("check your userName,there is no user %s", userName)
	}
	return findingUser, err
}
