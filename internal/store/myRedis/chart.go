package myRedis

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	"txnbi-backend/internal/model"
)

func SetExampleChart(ctx context.Context, v []model.ChartExample) error {
	key := "chart-example"
	vJSON, err := json.Marshal(v)
	if err != nil {
		return err
	}
	_, err = Cli.Set(ctx, key, vJSON, 0).Result()
	if err != nil {
		return err
	}

	return nil
}

func GetExampleChart(ctx context.Context) ([]model.ChartExample, error) {
	key := "chart-example"
	result, err := Cli.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	charts := make([]model.ChartExample, 0)
	err = json.Unmarshal([]byte(result), &charts)
	if err != nil {
		return nil, err
	}

	return charts, nil
}

func GetAllCharts(ctx context.Context, userID int64, offset int) (charts []model.Chart, total int64, err error) {
	// 先从缓存中获取数据
	key := fmt.Sprintf("user-allChart-%d:%d", offset, userID)
	keyCount := fmt.Sprintf("user-allChart-total:%d", userID)
	result, err := Cli.Get(ctx, key).Result()
	if err != nil {
		return nil, 0, err
	}
	resultTotal, err := Cli.Get(ctx, keyCount).Result()
	if err != nil {
		return nil, 0, err
	}

	// 将数据解析到 chart 切片中
	err = json.Unmarshal([]byte(result), &charts)
	if err != nil {
		return nil, 0, err
	}
	total, err = strconv.ParseInt(resultTotal, 10, 64)
	if err != nil {
		return nil, 0, err
	}
	return charts, total, nil
}

func SetAllCharts(ctx context.Context, userID, total int64, offset int, charts []model.Chart) error {
	key := fmt.Sprintf("user-allChart-%d:%d", offset, userID)
	keyCount := fmt.Sprintf("user-allChart-total:%d", userID)
	// 将结果缓存到 redis
	chartsJSON, err := json.Marshal(charts)
	if err != nil {
		return err
	}
	_, err = Cli.Set(ctx, key, chartsJSON, 24*time.Hour).Result()
	if err != nil {
		return err
	}
	_, err = Cli.Set(ctx, keyCount, total, 24*time.Hour).Result()
	if err != nil {
		return err
	}
	return nil
}
