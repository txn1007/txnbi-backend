package store

import (
	"context"
	"txnbi-backend/internal/model"
)

// AdminListCharts 管理员获取图表列表
func AdminListCharts(ctx context.Context, offset, limit int, keyword string, userID int64) ([]struct {
	model.Chart
	UserAccount string `gorm:"column:userAccount"`
}, int64, error) {
	var charts []struct {
		model.Chart
		UserAccount string `gorm:"column:userAccount"`
	}
	var total int64

	// 构建查询
	db := DB.Table("chart").
		Select("chart.*, users.userAccount as userAccount").
		Joins("left join users on chart.userId = users.id").
		Where("chart.isDelete = ?", 0)

	// 如果有关键字，添加搜索条件
	if keyword != "" {
		db = db.Where("chart.name LIKE ? OR chart.goal LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	}

	// 如果指定了用户ID，添加用户筛选条件
	if userID > 0 {
		db = db.Where("chart.userId = ?", userID)
	}

	// 获取总数
	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	err = db.Offset(offset).Limit(limit).Order("chart.updateTime DESC").Find(&charts).Error
	if err != nil {
		return nil, 0, err
	}

	return charts, total, nil
}

// GetChartDetailWithUser 获取图表详情（包含用户信息）
func GetChartDetailWithUser(ctx context.Context, chartID int64) (*struct {
	model.Chart
	UserAccount string `gorm:"column:userAccount"`
}, error) {
	var chart struct {
		model.Chart
		UserAccount string `gorm:"column:userAccount"`
	}

	err := DB.Table("chart").
		Select("chart.*, user.userAccount as userAccount").
		Joins("left join users on chart.userId = users.id").
		Where("chart.id = ? AND chart.isDelete = ?", chartID, 0).
		First(&chart).Error

	if err != nil {
		return nil, err
	}

	return &chart, nil
}

// DeleteChartByID 根据ID删除图表（软删除）
func DeleteUserChartByID(ctx context.Context, chartID int64, userID int64) error {
	return DB.Model(&model.Chart{}).
		Where("id = ? AND userId = ?", chartID, userID).
		Update("isDelete", 1).Error
}

// UpdateChart 更新图表信息
func UpdateChartInfo(ctx context.Context, chart *model.Chart) error {
	return DB.Model(&model.Chart{}).
		Where("id = ?", chart.ID).
		Updates(map[string]interface{}{
			"name":       chart.Name,
			"goal":       chart.Goal,
			"gen_result": chart.GenResult,
			"updateTime": chart.UpdateTime,
		}).Error
}

// GetChartsByUserID 获取用户的所有图表
func ListUserCharts(ctx context.Context, userID int64, page, pageSize int) ([]model.Chart, int64, error) {
	var charts []model.Chart
	var total int64

	offset := (page - 1) * pageSize

	// 获取总数
	err := DB.Model(&model.Chart{}).
		Where("userId = ? AND isDelete = ?", userID, 0).
		Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	err = DB.Where("userId = ? AND isDelete = ?", userID, 0).
		Offset(offset).
		Limit(pageSize).
		Order("updateTime DESC").
		Find(&charts).Error
	if err != nil {
		return nil, 0, err
	}

	return charts, total, nil
}

// BatchDeleteCharts 批量删除图表
func BatchDeleteCharts(ctx context.Context, chartIDs []int64) error {
	return DB.Model(&model.Chart{}).
		Where("id IN ?", chartIDs).
		Update("isDelete", 1).Error
}
