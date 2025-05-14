package biz

import (
	"context"
	"errors"
	"fmt"
	"txnbi-backend/api"
	"txnbi-backend/internal/store"

	"gorm.io/gorm"
)

// AdminListCharts 管理员获取图表列表
func AdminListCharts(ctx context.Context, page, pageSize int, keyword string, userID int64) ([]api.ChartInfoV1, int64, error) {
	offset := (page - 1) * pageSize
	charts, total, err := store.AdminListCharts(ctx, offset, pageSize, keyword, userID)
	if err != nil {
		return nil, 0, err
	}

	// 转换为响应格式
	result := make([]api.ChartInfoV1, 0, len(charts))
	for _, chart := range charts {
		result = append(result, api.ChartInfoV1{
			ChartID:     chart.ID,
			UserID:      chart.UserID,
			UserAccount: chart.UserAccount,
			ChartName:   chart.Name,
			ChartType:   chart.ChartType,
			ChartGoal:   chart.Goal,
			ChartCode:   chart.GenChart,
			ChartResult: chart.GenResult,
			Status:      chart.Status,
			UpdateTime:  chart.UpdateTime.Format("2006-01-02 15:04:05"),
		})
	}

	return result, total, nil
}

// AdminGetChartDetail 管理员获取图表详情
func AdminGetChartDetail(ctx context.Context, chartID int64) (*api.ChartInfoV1, error) {
	chart, err := store.GetChartDetailWithUser(ctx, chartID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("图表不存在")
		}
		return nil, err
	}

	return &api.ChartInfoV1{
		ChartID:     chart.ID,
		UserID:      chart.UserID,
		UserAccount: chart.UserAccount,
		ChartName:   chart.Name,
		ChartType:   chart.ChartType,
		ChartGoal:   chart.Goal,
		ChartCode:   chart.GenChart,
		ChartResult: chart.GenResult,
		Status:      chart.Status,
		UpdateTime:  chart.UpdateTime.Format("2006-01-02 15:04:05"),
	}, nil
}

// AdminUpdateChart 管理员更新图表
func AdminUpdateChart(ctx context.Context, chartID int64, chartName, goal, genResult string) error {
	chart, err := store.GetChartByID(ctx, chartID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("图表不存在")
		}
		return err
	}

	// 更新图表信息
	if chartName != "" {
		chart.Name = chartName
	}
	if goal != "" {
		chart.Goal = goal
	}
	if genResult != "" {
		chart.GenResult = genResult
	}

	return store.UpdateChart(ctx, chart)
}

// AdminDeleteChart 管理员删除图表
func AdminDeleteChart(ctx context.Context, chartID int64) error {
	chart, err := store.GetChartByID(ctx, chartID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("图表不存在")
		}
		return err
	}

	return store.DeleteChartByID(ctx, chartID, chart.UserID)
}
