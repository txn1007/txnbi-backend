package biz

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"txnbi-backend/internal/store"
)

func UserLogin(account string, password string) error {
	// 根据账号获取用户
	ac, err := store.GetUserByAccount(account)
	// 检查账号是否存在
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("account not exist")
	}
	// 检查密码是否正确
	if ac.UserPassword != password {
		return fmt.Errorf("password error")
	}
	// todo 生成并返回 token
	return nil
}

func UserRegister(account string, password string) error {
	ac, err := store.GetUserByAccount(account)
	if ac != nil && err == nil {
		return fmt.Errorf("account exist")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	err = store.CreateUser(account, password, "user")
	if err != nil {
		return err
	}
	return nil
}
