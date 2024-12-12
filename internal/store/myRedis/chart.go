package myRedis

import (
	"context"
	"encoding/json"
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
