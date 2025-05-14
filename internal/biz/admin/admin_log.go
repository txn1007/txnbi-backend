package biz

import (
	"context"
	"errors"
	"fmt"
	"time"
	"txnbi-backend/internal/model"
	"txnbi-backend/internal/store"

	"gorm.io/gorm"
)

// AdminListLogs 管理员获取日志列表
func AdminListLogs(ctx context.Context, page, pageSize int, keyword string, startTime, endTime time.Time) ([]model.OperationLog, int64, error) {
	offset := (page - 1) * pageSize
	return store.AdminListLogs(ctx, offset, pageSize, keyword, startTime, endTime)
}

// AdminGetLogDetail 管理员获取日志详情
func AdminGetLogDetail(ctx context.Context, logID int64) (*model.OperationLog, error) {
	log, err := store.AdminGetLogByID(ctx, logID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("日志不存在")
		}
		return nil, err
	}
	return log, nil
}

// AdminCreateLog 管理员创建日志
func AdminCreateLog(ctx context.Context, log *model.OperationLog) (int64, error) {
	// 获取用户信息
	user, err := store.GetUserByID(log.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, fmt.Errorf("用户不存在")
		}
		return 0, err
	}

	// 设置用户名和账号
	log.UserName = user.UserName
	log.UserAccount = user.UserAccount
	log.CreateTime = time.Now()

	return store.AdminCreateLog(ctx, log)
}

// AdminUpdateLog 管理员更新日志
func AdminUpdateLog(ctx context.Context, log *model.OperationLog) error {
	// 检查日志是否存在
	existLog, err := store.AdminGetLogByID(ctx, log.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("日志不存在")
		}
		return err
	}

	// 保留不更新的字段
	log.UserID = existLog.UserID
	log.UserName = existLog.UserName
	log.UserAccount = existLog.UserAccount
	log.CreateTime = existLog.CreateTime

	return store.AdminUpdateLog(ctx, log)
}

// AdminDeleteLog 管理员删除日志
func AdminDeleteLog(ctx context.Context, logID int64) error {
	// 检查日志是否存在
	_, err := store.AdminGetLogByID(ctx, logID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("日志不存在")
		}
		return err
	}

	return store.AdminDeleteLog(ctx, logID)
}

// AdminBatchDeleteLogs 管理员批量删除日志
func AdminBatchDeleteLogs(ctx context.Context, logIDs []int64) error {
	return store.AdminBatchDeleteLogs(ctx, logIDs)
}
