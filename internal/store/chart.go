package store

import (
	"txnbi-backend/internal/model"
)

func CreateChart(chartName string, goal, chartData, genChart, genResult, chartType string, userID int64) error {
	return DB.Create(&model.Chart{Name: chartName, Goal: goal, ChartData: chartData, ChartType: chartType,
		UserID: userID, GenChart: genChart, GenResult: genResult}).Error
}

// FindChartAndPage 如果没有输入表名，则返回所有记录
// 如果输入，则查找表名对应的记录
// 分页查询
func FindChartAndPage(userID int64, chartName string, currentPage, pageSize int) (charts []model.Chart, total int64, err error) {
	// 计算偏移量
	offset := (currentPage - 1) * pageSize

	// 查询所有记录
	charts = make([]model.Chart, 0, pageSize)
	if chartName == "" {
		err = DB.Offset(offset).Limit(pageSize).Where("userId = ?", userID).Find(&charts).Error
		if err != nil {
			return nil, 0, err
		}
		err = DB.Model(&model.Chart{}).Where("userId = ?", userID).Count(&total).Error
		if err != nil {
			return nil, 0, err
		}
	} else {
		err = DB.Offset(offset).Limit(pageSize).Where("userId = ? AND name = ?", userID, chartName).Find(&charts).Error
		if err != nil {
			return nil, 0, err
		}
		err = DB.Model(&model.Chart{}).Where("userId = ? AND name = ?", userID, chartName).Count(&total).Error
		if err != nil {
			return nil, 0, err
		}
	}

	return charts, total, nil
}
