package biz

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/xuri/excelize/v2"
	"mime/multipart"
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

	// AI生成数据
	chartData, analysis, err = doubao.GenChart(goal, csvStr, chartType)
	if err != nil {
		return "", "", err
	}

	// 存入数据库
	err = store.CreateChart(chartName, goal, csvStr, chartData, analysis, chartType, userID)
	if err != nil {
		return "", "", err
	}

	return chartData, analysis, nil
}
