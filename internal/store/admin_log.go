package store

import (
	"context"
	"time"
	"txnbi-backend/internal/model"
)

// AdminListLogs 管理员获取日志列表
func AdminListLogs(ctx context.Context, offset, limit int, keyword string, startTime, endTime time.Time) ([]model.OperationLog, int64, error) {
	var logs []model.OperationLog
	var total int64

	db := DB.Model(&model.OperationLog{})

	// 添加时间范围条件
	if !startTime.IsZero() && !endTime.IsZero() {
		db = db.Where("create_time BETWEEN ? AND ?", startTime, endTime)
	}

	// 如果有关键字，添加搜索条件
	if keyword != "" {
		db = db.Where("operation LIKE ? OR user_name LIKE ? OR user_account LIKE ?",
			"%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")
	}

	// 获取总数
	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	err = db.Offset(offset).Limit(limit).Order("create_time DESC").Find(&logs).Error
	if err != nil {
		return nil, 0, err
	}

	return logs, total, nil
}

// AdminGetLogByID 管理员根据ID获取日志
func AdminGetLogByID(ctx context.Context, logID int64) (*model.OperationLog, error) {
	var log model.OperationLog
	err := DB.Where("id = ?", logID).First(&log).Error
	if err != nil {
		return nil, err
	}
	return &log, nil
}

// AdminCreateLog 管理员创建日志
func AdminCreateLog(ctx context.Context, log *model.OperationLog) (int64, error) {
	err := DB.Create(log).Error
	if err != nil {
		return 0, err
	}
	return int64(log.ID), nil
}

// AdminUpdateLog 管理员更新日志
func AdminUpdateLog(ctx context.Context, log *model.OperationLog) error {
	return DB.Model(&model.OperationLog{}).Where("id = ?", log.ID).Updates(map[string]interface{}{
		"operation": log.Operation,
		"method":    log.Method,
		"path":      log.Path,
		"ip":        log.IP,
	}).Error
}

// AdminDeleteLog 管理员删除日志
func AdminDeleteLog(ctx context.Context, logID int64) error {
	return DB.Delete(&model.OperationLog{}, logID).Error
}

// AdminBatchDeleteLogs 管理员批量删除日志
func AdminBatchDeleteLogs(ctx context.Context, logIDs []int64) error {
	return DB.Where("id IN ?", logIDs).Delete(&model.OperationLog{}).Error
}
