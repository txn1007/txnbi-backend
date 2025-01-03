package biz

import (
	"bytes"
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"mime/multipart"
	"time"
	"txnbi-backend/api"
	"txnbi-backend/conf"
	"txnbi-backend/errs"
	"txnbi-backend/internal/model"
	"txnbi-backend/internal/store"
	"txnbi-backend/pkg/doubao"
	"txnbi-backend/pkg/jwt"
)

func GenChart(ctx context.Context, chartName, chartType, goal string, data *multipart.FileHeader, userID int64) (chartData, analysis string, err error) {
	// 将上传的excel数据转为csv类型字符串
	fd, err := data.Open()
	if err != nil {
		return "", "", err
	}
	defer fd.Close()
	reader, err := excelize.OpenReader(fd)
	if err != nil {
		return "", "", err
	}
	rows, err := reader.GetRows("Sheet1")
	if err != nil {
		return "", "", err
	}
	// 写入 CSV
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)
	for _, row := range rows {
		if err = writer.Write(row); err != nil {
			return "", "", fmt.Errorf("写入 CSV 失败: %v", err)
		}
	}
	writer.Flush()
	if writer.Error() != nil {
		return "", "", writer.Error()
	}
	csvStr := buf.String()

	chartID := time.Now().Unix()
	// 生成用户生成的图表，其数据库表结构
	DBChartName, err := store.CreateUserGenChart(chartID, rows)
	if err != nil {
		return "", "", err
	}

	// AI生成数据
	chartData, analysis, err = doubao.GenChart(goal, csvStr, chartType)
	if err != nil {
		return "", "", err
	}

	// 存入数据库
	chart := model.Chart{Name: chartName, UserID: userID, ChartTableName: DBChartName, Goal: goal, GenChart: chartData, GenResult: analysis, ChartType: chartType}
	err = store.CreateChart(ctx, chart)
	if err != nil {
		return "", "", err
	}

	return chartData, analysis, nil
}

func ListMyChart(ctx context.Context, userID int64, chartName string, currentPage int, pageSize int) (charts []api.ChartInfoV0, total int64, err error) {
	var modelCharts []model.Chart
	// 根据是否查询表名参数，查询数据库
	if chartName == "" {
		modelCharts, total, err = store.GetChartsByUserID(ctx, userID, currentPage, pageSize)
		if err != nil {
			return nil, 0, err
		}
	} else {
		modelCharts, total, err = store.FindChartsByUserIDAndChartNane(ctx, userID, chartName, currentPage, pageSize)
		if err != nil {
			return nil, 0, err
		}
	}
	// 将 model.Chart 转化为 api.ChartInfoV0
	charts = make([]api.ChartInfoV0, len(modelCharts))
	for i, chart := range modelCharts {
		charts[i] = api.ChartInfoV0{
			ChartID:     chart.ID,
			ChartType:   chart.ChartType,
			ChartGoal:   chart.Goal,
			ChartName:   chart.Name,
			ChartCode:   chart.GenChart,
			ChartResult: chart.GenResult,
			UpdateTime:  chart.UpdateTime.Format("2006-01-02 15:04:05"),
		}
	}
	return charts, total, nil
}

func DeleteMyChart(ctx context.Context, chartID, userID int64) error {
	// 获取图表信息
	chart, err := store.GetChartByID(ctx, chartID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("图表不存在！")
	}
	// 检查是否为请求发起者创建的图表
	if chart.UserID != userID {
		return fmt.Errorf("该图表并非由您创建！")
	}
	return store.DeleteChartByID(ctx, chartID, userID)
}

func ExampleChart(ctx context.Context) ([]api.ChartInfoV0, int64, error) {
	//从缓存中获取展示表
	charts, total, err := store.GetExampleChartByRedis(ctx)
	if err != nil {
		return nil, 0, err
	}
	return charts, total, nil

	//// 获取本地硬编码的示例数据
	//return store.GetExampleChartByLocal(ctx)
}

func UpdateChart(ctx context.Context, chartID, userID int64, chartName, goal, genResult string) error {
	chart, err := store.GetChartByID(ctx, chartID)
	// 要修改的表不存在
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errs.ErrFindNotExistChart
	}
	// 未知错误
	if err != nil {
		return err
	}
	// 比较请求修改的表是否属于该用户
	if chart.UserID != userID {
		return errs.ErrOperateOtherUserChart
	}

	// 修改记录
	chart.Name = chartName
	chart.Goal = goal
	chart.GenResult = genResult
	err = store.UpdateChart(ctx, chart)
	if err != nil {
		return err
	}

	return nil
}

func ShareChart(ctx context.Context, chartID, userID int64, userToken string) (accessCode string, err error) {
	chart, err := store.GetChartByID(ctx, chartID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", errs.ErrFindNotExistChart
	}
	if err != nil {
		return "", err
	}

	// 检查表是否属于用户
	if chart.UserID != userID {
		return "", errs.ErrOperateOtherUserChart
	}

	// 生成短期 token 用作 accessCode 以访问图表。当用户token过期时，accessCode 失效
	accessCode, err = jwt.SignForChartAccessCode(userID, chartID, userToken, conf.JWTCfg.SignKey)
	if err != nil {
		return "", err
	}
	return accessCode, nil
}
