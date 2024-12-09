package store

import (
	"fmt"
	"regexp"
	"strings"
	"txnbi-backend/internal/model"
)

func CreateChart(chartName, chartTableName, goal, genChart, genResult, chartType string, userID int64) error {
	return DB.Create(&model.Chart{Name: chartName, Goal: goal, ChartTableName: chartTableName, ChartType: chartType,
		UserID: userID, GenChart: genChart, GenResult: genResult}).Error
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

// FindChartAndPage 如果没有输入表名，则返回所有记录
// 如果输入，则查找表名对应的记录
// 分页查询
func FindChartAndPage(userID int64, chartName string, currentPage, pageSize int) (charts []model.Chart, total int64, err error) {
	// 计算偏移量
	offset := (currentPage - 1) * pageSize

	// 查询所有记录
	charts = make([]model.Chart, 0, pageSize)
	// 根据是否为查询构建 SQL的where语句
	var whereSQL string
	if chartName != "" {
		whereSQL = fmt.Sprintf("name = '%s' AND userId = %d", chartName, userID)
	} else {
		whereSQL = fmt.Sprintf("userId = %d", userID)
	}

	// 查询记录与总数
	err = DB.Offset(offset).Limit(pageSize).
		Select("id", "chartType", "name", "goal", "genChart", "genResult", "updateTime").
		Where(whereSQL).Order("updateTime desc").Find(&charts).Error
	if err != nil {
		return nil, 0, err
	}
	err = DB.Model(&model.Chart{}).Where(whereSQL).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	return charts, total, nil
}
