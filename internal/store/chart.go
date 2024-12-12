package store

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"regexp"
	"strings"
	"time"
	"txnbi-backend/api"
	"txnbi-backend/internal/model"
	"txnbi-backend/internal/store/myRedis"
)

func CreateChart(ctx context.Context, chart model.Chart) error {
	err := DB.Create(&chart).Error

	// 删除创建该图表的用户的，查看图表所有页的缓存
	// 删除总页数
	_, err = myRedis.Cli.Del(ctx, fmt.Sprintf("user-allChart-total:%d", chart.UserID)).Result()
	if err != nil {
		return err
	}
	// 删除查看图表的分页数据缓存
	pageKeys, err := myRedis.Cli.Keys(ctx, "user-allChart-*").Result()
	if err != nil {
		return err
	}
	for _, pageKey := range pageKeys {
		err = myRedis.Cli.Del(ctx, pageKey).Err()
		if err != nil {
			return err
		}
	}

	return nil
}

func CreateUserGenChart(chartID int64, excelRow [][]string) (DBTableName string, err error) {
	// 定义表名
	tableName := fmt.Sprintf("userGenChart_%d", chartID)

	// 1.生成表
	// 构造建表语句的字段SQL部分
	var columns []string
	fields := excelRow[0]
	for _, field := range fields {
		// 检查字段名是否合法
		// 字段名只允许汉字、英文、数字、下划线
		matched, err := regexp.MatchString(`^[\p{Han}_a-zA-Z0-9]+$`, field)
		if err != nil {
			return "", err
		}
		if !matched {
			return "", fmt.Errorf("字段命名非法")
		}

		columns = append(columns, fmt.Sprintf("`%s` VARCHAR(256) NOT NULL", field))
	}
	createStmt := fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s` (%s) collate = utf8mb4_unicode_ci;", tableName, strings.Join(columns, ", "))

	// 执行 CREATE TABLE 语句
	if err = DB.Exec(createStmt).Error; err != nil {
		return "", err
	}

	// 2.参数化插入用户数据
	var valsSQL, vals []string
	for i := 1; i < len(excelRow); i++ {
		curVal := excelRow[i]
		// 将数据插入到建表语句中
		inner := strings.Repeat("? ,", len(curVal))
		curValSQL := fmt.Sprintf("(%s)", inner[:len(inner)-1])
		valsSQL = append(valsSQL, curValSQL)
		// 添加所有数据值到vals
		vals = append(vals, curVal...)
	}

	// 预编译插入语句，防止SQL注入
	db, err := DB.DB()
	if err != nil {
		return "", err
	}
	insertStmt, err := db.Prepare(fmt.Sprintf("INSERT INTO %s VALUES %s", tableName, strings.Join(valsSQL, ",")))
	if err != nil {
		return "", err
	}

	// 执行 INSERT 语句
	// 手动转换 []string 到 []any
	anyVals := make([]any, len(vals))
	for i, v := range vals {
		anyVals[i] = v
	}
	if _, err = insertStmt.Exec(anyVals...); err != nil {
		return "", err
	}
	return tableName, nil
}

// GetChartsByUserID 分页获取该用户的图表信息
func GetChartsByUserID(ctx context.Context, userID int64, currentPage, pageSize int) (charts []model.Chart, total int64, err error) {
	charts = make([]model.Chart, 0, pageSize)

	offset := (currentPage - 1) * pageSize
	// 先从缓存中获取数据
	charts, total, err = myRedis.GetAllCharts(ctx, userID, offset)
	// 缓存不存在则从 DB 层获取数据
	if err != nil && errors.Is(err, redis.Nil) {
		whereSQL := fmt.Sprintf("userId = %d AND isDelete = 0", userID)
		// 从 DB 层获取数据
		err = DB.Offset(offset).Limit(pageSize).
			Select("id", "chartType", "name", "goal", "genChart", "genResult", "updateTime").
			Where(whereSQL + " AND isDelete = 0").Order("updateTime desc").Find(&charts).Error
		if err != nil {
			return nil, 0, err
		}
		err = DB.Model(&model.Chart{}).Where(whereSQL + " AND isDelete = 0").Count(&total).Error
		if err != nil {
			return nil, 0, err
		}
		// 将结果缓存到 redis
		err = myRedis.SetAllCharts(ctx, userID, total, offset, charts)
		if err != nil {
			return nil, 0, err
		}
		return charts, total, nil
	}
	// 缓存中存在，则将缓存中的数据解析到 chart切片中
	err = myRedis.SetAllCharts(ctx, userID, total, offset, charts)
	if err != nil {
		return nil, 0, err
	}
	return charts, total, nil
}

// FindChartsByUserIDAndChartNane 分页查找该用户的图表信息
func FindChartsByUserIDAndChartNane(ctx context.Context, userID int64, chartName string, currentPage, pageSize int) (charts []model.Chart, total int64, err error) {
	charts = make([]model.Chart, 0, pageSize)

	offset := (currentPage - 1) * pageSize
	whereSQL := fmt.Sprintf("name = '%s' AND userId = %d AND isDelete = 0", chartName, userID)

	// 搜索图表
	err = DB.Offset(offset).Limit(pageSize).
		Select("id", "chartType", "name", "goal", "genChart", "genResult", "updateTime").
		Where(whereSQL).Order("updateTime desc").Find(&charts).Error
	if err != nil {
		return nil, 0, err
	}

	// 计算总数
	err = DB.Model(&model.Chart{}).Where(whereSQL).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	return charts, total, nil
}

// DeleteChartByID 软删除
func DeleteChartByID(ctx context.Context, chartID, userID int64) error {
	// 先软删除数据库的数据，再删除 redis 中的缓存
	key := fmt.Sprintf("chart-id:%d", chartID)

	err := DB.Model(&model.Chart{}).Where("id = ?", chartID).Update("isDelete", 1).Error
	if err != nil {
		return err
	}
	// 删除缓存
	_, err = myRedis.Cli.Del(ctx, key).Result()
	if err != nil {
		return err
	}

	// 删除创建该图表的用户的，查看图表所有页的缓存
	// 删除总页数
	_, err = myRedis.Cli.Del(ctx, fmt.Sprintf("user-allChart-total:%d", userID)).Result()
	if err != nil {
		return err
	}
	// 删除查看图表的分页数据缓存
	pageKeys, err := myRedis.Cli.Keys(ctx, "user-allChart-*").Result()
	if err != nil {
		return err
	}
	for _, pageKey := range pageKeys {
		err = myRedis.Cli.Del(ctx, pageKey).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

func GetChartByID(ctx context.Context, chartID int64) (*model.Chart, error) {
	var chart model.Chart
	// 从redis中获取
	key := fmt.Sprintf("chart-id:%d", chartID)
	result, err := myRedis.Cli.Get(ctx, key).Result()

	// 如果 redis 中不存在，则从 DB 层获取数据，并缓存到 redis 中
	if errors.Is(err, redis.Nil) {
		err = DB.Where("id = ? AND isDelete = 0", chartID).First(&chart).Error
		if err != nil {
			return nil, err
		}
		// 缓存到redis
		chartIDJSON, err := json.Marshal(chart)
		if err != nil {
			return nil, err
		}
		_, err = myRedis.Cli.Set(ctx, key, chartIDJSON, 3*24*time.Hour).Result()
		if err != nil {
			return nil, err
		}
		return &chart, nil
	}

	// 如果redis中存在,则反序列化
	err = json.Unmarshal([]byte(result), &chart)
	if err != nil {
		return nil, err
	}

	return &chart, nil
}

func GetExampleChartByRedis(ctx context.Context) (apiCharts []api.ChartInfoV0, total int64, err error) {
	// 从 Redis 中获取数据
	charts, err := myRedis.GetExampleChart(ctx)
	// 只能从缓存中获取
	if errors.Is(err, redis.Nil) {
		return nil, 0, fmt.Errorf("在缓存中查询不到示例数据！")
	}
	if err != nil {
		return nil, 0, err
	}

	// 转化
	apiCharts = make([]api.ChartInfoV0, len(charts))
	for i, chart := range charts {
		apiCharts[i] = api.ChartInfoV0{
			ChartID:     chart.ID,
			ChartType:   chart.ChartType,
			ChartGoal:   chart.Goal,
			ChartName:   chart.Name,
			ChartCode:   chart.GenChart,
			ChartResult: chart.GenResult,
			UpdateTime:  chart.UpdateTime.Format("2006-01-02 15:04:05"),
		}
	}
	return apiCharts, int64(len(charts)), nil
}

func GetExampleChartByLocal(ctx context.Context) (charts []api.ChartInfoV0, total int64, err error) {
	return ExampleChart, int64(len(ExampleChart)), nil
}
