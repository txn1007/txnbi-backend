package store

import (
	"fmt"
	"strings"
	"txnbi-backend/internal/model"
)

func CreateChart(chartName, chartTableName, goal, genChart, genResult, chartType string, userID int64) error {
	return DB.Create(&model.Chart{Name: chartName, Goal: goal, ChartTableName: chartTableName, ChartType: chartType,
		UserID: userID, GenChart: genChart, GenResult: genResult}).Error
}

func CreateUserGenChart(chartID int64, excelRow [][]string) (DBTableName string, err error) {
	// 构建建表语句并执行
	// 定义表名
	tableName := fmt.Sprintf("userGenChart_%d", chartID)
	// 构造建表语句的字段SQL部分
	var columns []string
	fields := excelRow[0]
	for _, field := range fields {
		column := fmt.Sprintf("`%s` VARCHAR(256) NOT NULL", field)
		columns = append(columns, column)
	}
	createStmt := fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s` (%s) collate = utf8mb4_unicode_ci;", tableName, strings.Join(columns, ", "))
	// 执行 CREATE TABLE 语句
	if err = DB.Exec(createStmt).Error; err != nil {
		return "", err
	}

	// 将数据插入到建表语句中
	var vals []string
	for i := 1; i < len(excelRow); i++ {
		curVal := excelRow[i]
		curValSQL := fmt.Sprintf("(%s)", strings.Join(curVal, ","))
		vals = append(vals, curValSQL)
	}
	insertStmt := fmt.Sprintf("INSERT INTO %s VALUES %s", tableName, strings.Join(vals, ","))
	// 执行 INSERT 语句
	if err = DB.Exec(insertStmt).Error; err != nil {
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
