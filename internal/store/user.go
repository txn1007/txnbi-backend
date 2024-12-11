package store

import (
	"errors"
	"fmt"
	"time"
	"txnbi-backend/internal/model"
)

func GetUserByAccount(account string) (*model.User, error) {
	var u model.User
	err := DB.Where("userAccount = ?", account).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func GetUserByID(id int64) (*model.User, error) {
	var u model.User
	err := DB.Where("id = ?", id).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func CreateUser(u model.User) (userID int64, err error) {
	// 非法角色
	if !(u.UserRole == "admin" || u.UserRole == "user") {
		return 0, errors.New("user role must be admin or user")
	}
	randomUserName := fmt.Sprintf("user_%d", time.Now().Unix()%1000000)
	u.UserName = randomUserName
	u.UserAvatar = "https://tiktokk-1331222828.cos.ap-guangzhou.myqcloud.com/avatar/avatar-tem.jpg"
	err = DB.Create(&u).Error
	if err != nil {
		return 0, err
	}
	return u.ID, nil
}
