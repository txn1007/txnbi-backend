package store

import "txnbi-backend/internal/model"

func CreateChart(chartName string, goal, chartData, genChart, genResult, chartType string, userID int64) error {
	return DB.Create(&model.Chart{Name: chartName, Goal: goal, ChartData: chartData, ChartType: chartType,
		UserID: userID, GenChart: genChart, GenResult: genResult}).Error
}
