package store

import (
	"context"
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
	return int64(u.ID), nil
}

// 在 store 包中添加以下方法

// GetUserList 获取用户列表
func GetUserList(ctx context.Context, page, pageSize int, keyword string) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	db := DB.Model(&model.User{})

	// 如果有关键字，添加搜索条件
	if keyword != "" {
		db = db.Where("user_account LIKE ?", "%"+keyword+"%")
	}

	// 获取总数
	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	err = db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// GetChartList 获取图表列表
func GetChartList(ctx context.Context, page, pageSize int, keyword string) ([]model.Chart, int64, error) {
	var charts []model.Chart
	var total int64

	db := DB.Model(&model.Chart{}).Where("is_delete = ?", 0)

	// 如果有关键字，添加搜索条件
	if keyword != "" {
		db = db.Where("name LIKE ? OR goal LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}

	// 获取总数
	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	err = db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&charts).Error
	if err != nil {
		return nil, 0, err
	}

	return charts, total, nil
}

// DeleteChart 删除图表
func DeleteChart(ctx context.Context, chartID int64) error {
	// 软删除
	return DB.Model(&model.Chart{}).Where("id = ?", chartID).Update("is_delete", 1).Error
}

// GetLogList 获取日志列表
func GetLogList(ctx context.Context, page, pageSize int, keyword string, startTime, endTime time.Time) ([]model.OperationLog, int64, error) {
	var logs []model.OperationLog
	var total int64

	db := DB.Model(&model.OperationLog{})

	// 添加时间范围条件
	if !startTime.IsZero() && !endTime.IsZero() {
		db = db.Where("create_time BETWEEN ? AND ?", startTime, endTime)
	}

	// 如果有关键字，添加搜索条件
	if keyword != "" {
		db = db.Where("operation LIKE ? OR user_name LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}

	// 获取总数
	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	err = db.Offset((page - 1) * pageSize).Limit(pageSize).Order("create_time DESC").Find(&logs).Error
	if err != nil {
		return nil, 0, err
	}

	return logs, total, nil
}

// DeleteLog 删除日志
func DeleteLog(ctx context.Context, logID int64) error {
	return DB.Delete(&model.OperationLog{}, logID).Error
}

// DeleteLogBatch 批量删除日志
func DeleteLogBatch(ctx context.Context, logIDs []int64) error {
	return DB.Delete(&model.OperationLog{}, logIDs).Error
}

// GetRoleList 获取角色列表
func GetRoleList(ctx context.Context) ([]model.Role, error) {
	var roles []model.Role
	err := DB.Find(&roles).Error
	return roles, err
}

// GetRoleByID 根据ID获取角色
func GetRoleByID(ctx context.Context, roleID int64) (*model.Role, error) {
	var role model.Role
	err := DB.Where("id = ?", roleID).First(&role).Error
	if err != nil {
		return nil, err
	}

	return &role, nil
}

// GetPermissionList 获取权限列表
func GetPermissionList(ctx context.Context) ([]model.Permission, error) {
	var permissions []model.Permission
	err := DB.Find(&permissions).Error
	return permissions, err
}

// AssignRoleToUser 为用户分配角色
func AssignRoleToUser(ctx context.Context, userID int64, roleID int64) error {
	userRole := model.UserRole{
		UserID: uint64(userID),
		RoleID: uint64(roleID),
	}

	return DB.Create(&userRole).Error
}

// UpdateUserRole 更新用户角色
func UpdateUserRole(ctx context.Context, userID int64, roleID int64) error {
	// 先删除用户的所有角色
	err := DB.Where("user_id = ?", userID).Delete(&model.UserRole{}).Error
	if err != nil {
		return err
	}

	// 分配新角色
	return AssignRoleToUser(ctx, userID, roleID)
}
