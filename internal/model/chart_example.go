// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameChartExample = "chart_example"

// ChartExample 示例图表表
type ChartExample struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:id" json:"id"`                          // id
	Goal       string    `gorm:"column:goal;comment:分析目标" json:"goal"`                                                  // 分析目标
	Name       string    `gorm:"column:name;comment:图表名称" json:"name"`                                                  // 图表名称
	ChartType  string    `gorm:"column:chartType;comment:图表类型" json:"chartType"`                                        // 图表类型
	GenChart   string    `gorm:"column:genChart;comment:生成的图表数据" json:"genChart"`                                       // 生成的图表数据
	GenResult  string    `gorm:"column:genResult;comment:生成的分析结论" json:"genResult"`                                     // 生成的分析结论
	Status     string    `gorm:"column:status;not null;default:wait;comment:wait,running,succeed,failed" json:"status"` // wait,running,succeed,failed
	CreateTime time.Time `gorm:"column:createTime;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"createTime"`   // 创建时间
	UpdateTime time.Time `gorm:"column:updateTime;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updateTime"`   // 更新时间
}

// TableName ChartExample's table name
func (*ChartExample) TableName() string {
	return TableNameChartExample
}
