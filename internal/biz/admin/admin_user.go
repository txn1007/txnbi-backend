package biz

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"txnbi-backend/internal/model"
	"txnbi-backend/internal/store"
	"txnbi-backend/tool/encry"
)

// ListUsers 获取用户列表
func ListUsers(ctx context.Context, page, pageSize int) ([]*model.User, int64, error) {
	offset := (page - 1) * pageSize
	return store.ListUsers(ctx, offset, pageSize)
}

// GetUserDetail 获取用户详情
func GetUserDetail(ctx context.Context, userID int64) (*model.User, error) {
	return store.GetUserByID(userID)
}

// UpdateUser 更新用户信息
func UpdateUser(ctx context.Context, user *model.User) error {
	// 检查用户是否存在
	existUser, err := store.GetUserByID(user.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("用户不存在")
		}
		return err
	}
	
	// 如果密码有更新，需要加密
	if user.UserPassword != "" && user.UserPassword != existUser.UserPassword {
		user.UserPassword = encry.EncodeByMd5(user.UserPassword)
	} else {
		user.UserPassword = existUser.UserPassword
	}
	
	return store.UpdateUser(ctx, user)
}

// DeleteUser 删除用户
func DeleteUser(ctx context.Context, userID int64) error {
	// 检查用户是否存在
	_, err := store.GetUserByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("用户不存在")
		}
		return err
	}
	
	return store.DeleteUser(ctx, userID)
}

// DisableUser 禁用/启用用户
func DisableUser(ctx context.Context, userID int64, status int) error {
	// 检查用户是否存在
	_, err := store.GetUserByID(userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("用户不存在")
		}
		return err
	}
	
	return store.DisableUser(ctx, userID, status)
}

// SearchUsers 搜索用户
func SearchUsers(ctx context.Context, keyword string, page, pageSize int) ([]*model.User, int64, error) {
	offset := (page - 1) * pageSize
	return store.SearchUsers(ctx, keyword, offset, pageSize)
}

// CreateUser 创建用户（管理员创建）
func CreateUser(ctx context.Context, account, password, role string) (int64, error) {
	// 检查账号是否已存在
	existUser, err := store.GetUserByAccount(account)
	if existUser != nil && err == nil {
		return 0, fmt.Errorf("账号已存在")
	}
	
	// 创建用户
	user := model.User{
		UserAccount:  account,
		UserPassword: encry.EncodeByMd5(password),
		UserRole:     role,
		UserStatus:   0, // 默认正常状态
	}
	
	return store.CreateUser(user)
}