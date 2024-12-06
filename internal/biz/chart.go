package biz

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/xuri/excelize/v2"
	"mime/multipart"
	"time"
	"txnbi-backend/api"
	"txnbi-backend/internal/store"
	"txnbi-backend/pkg/doubao"
)

func GenChart(chartName, chartType, goal string, data *multipart.FileHeader, userID int64) (chartData, analysis string, err error) {
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
	err = store.CreateChart(chartName, DBChartName, goal, chartData, analysis, chartType, userID)
	if err != nil {
		return "", "", err
	}

	return chartData, analysis, nil
}

func ListMyChart(userID int64, chartName string, currentPage int, pageSize int) ([]api.ChartInfoV0, int64, error) {
	// 查询数据库
	charts, total, err := store.FindChartAndPage(userID, chartName, currentPage, pageSize)
	if err != nil {
		return nil, 0, err
	}
	// 转化
	apiCharts := make([]api.ChartInfoV0, len(charts))
	for i, chart := range charts {
		apiCharts[i] = api.ChartInfoV0{
			ChartID:     chart.ID,
			ChartType:   chart.ChartType,
			ChartGoal:   chart.Goal,
			ChartName:   chart.Name,
			ChartCode:   chart.GenChart,
			ChartResult: chart.GenResult,
		}
	}
	return apiCharts, total, nil
}
