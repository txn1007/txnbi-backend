package biz

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"txnbi-backend/conf"
	"txnbi-backend/internal/model"
	"txnbi-backend/internal/store"
	"txnbi-backend/pkg/jwt"
	"txnbi-backend/pkg/myRedis"
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
	// 将token记录到redis中，标记为最新的token
	token = jwt.SignForUser(ac.ID, ac.UserAccount, conf.JWTCfg.SignKey)
	myRedis.SetUserToken(ac.ID, token)
	return token, nil
}

func UserRegister(account, password, inviteCode string) error {
	// 检查邀请码是否存在
	is, err := myRedis.IsInviteCode(inviteCode)
	if !is {
		return fmt.Errorf("邀请码不存在！")
	}
	if err != nil {
		return err
	}
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

func UserLoginOut(userID int64) error {
	myRedis.DeleteUserToken(userID)
	return nil
}
