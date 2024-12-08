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

func CreateUser(account string, password string, userRole string) error {
	// 非法角色
	if !(userRole == "admin" || userRole == "user") {
		return errors.New("user role must be admin or user")
	}
	randomUserName := fmt.Sprintf("user_%d", time.Now().Unix()%1000000)
	return DB.Create(&model.User{UserAccount: account, UserPassword: password, UserName: randomUserName, UserRole: userRole, UserAvatar: "https://tiktokk-1331222828.cos.ap-guangzhou.myqcloud.com/avatar/avatar-tem.jpg"}).Error
}
