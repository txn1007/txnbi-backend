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
