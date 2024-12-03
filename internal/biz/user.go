package biz

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"txnbi-backend/conf"
	"txnbi-backend/internal/model"
	"txnbi-backend/internal/store"
	"txnbi-backend/pkg/jwt"
	"txnbi-backend/tool/encry"
)

func UserLogin(account string, password string) (token string, err error) {
	// 根据账号获取用户
	ac, err := store.GetUserByAccount(account)
	// 检查账号是否存在
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", fmt.Errorf("account not exist")
	}
	// 检查密码是否正确
	if ac.UserPassword != encry.EncodeByMd5(password) {
		return "", fmt.Errorf("password error")
	}
	return jwt.SignForUser(ac.ID, ac.UserAccount, conf.JWTCfg.SignKey), nil
}

func UserRegister(account string, password string) error {
	ac, err := store.GetUserByAccount(account)
	if ac != nil && err == nil {
		return fmt.Errorf("account exist")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// 加密用户密码
	err = store.CreateUser(account, encry.EncodeByMd5(password), "user")
	if err != nil {
		return err
	}
	return nil
}

func CurrentUserDetail(userID int64) (*model.User, error) {
	return store.GetUserByID(userID)
}
