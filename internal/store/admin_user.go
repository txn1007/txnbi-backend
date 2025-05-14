package store

import (
	"context"
	"time"
	"txnbi-backend/internal/model"
)

// ListUsers 获取用户列表
func ListUsers(ctx context.Context, offset, limit int) (users []*model.User, total int64, err error) {
	users = make([]*model.User, 0)
	// 查询总数
	if err = DB.Model(&model.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	// 分页查询
	if err = DB.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, 0, err
	}
	return users, total, nil
}

// UpdateUser 更新用户信息
func UpdateUser(ctx context.Context, user *model.User) error {
	user.UpdateTime = time.Now()
	return DB.Model(&model.User{}).Where("id = ?", user.ID).Updates(user).Error
}

// DeleteUser 删除用户
func DeleteUser(ctx context.Context, userID int64) error {
	return DB.Delete(&model.User{}, userID).Error
}

// DisableUser 禁用用户
func DisableUser(ctx context.Context, userID int64, status int) error {
	return DB.Model(&model.User{}).Where("id = ?", userID).Update("userStatus", status).Error
}

// SearchUsers 搜索用户
func SearchUsers(ctx context.Context, keyword string, offset, limit int) (users []*model.User, total int64, err error) {
	users = make([]*model.User, 0)
	db := DB.Model(&model.User{}).Where("userAccount LIKE ? OR userName LIKE ?", "%"+keyword+"%", "%"+keyword+"%")

	// 查询总数
	if err = db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	if err = db.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}
