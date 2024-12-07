package model

import (
	"time"
)

const TableNameChart = "chart"

// Chart 图表信息表
type Chart struct {
	ID             int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:id" json:"id"`                          // id
	Goal           string    `gorm:"column:goal;comment:分析目标" json:"goal"`                                                  // 分析目标
	Name           string    `gorm:"column:name;comment:图表名称" json:"name"`                                                  // 图表名称
	ChartTableName string    `gorm:"column:chartTableName;comment:用户上传的图表的MySQL表名称" json:"chartTableName"`                  // 图表数据
	ChartType      string    `gorm:"column:chartType;comment:图表类型" json:"chartType"`                                        // 图表类型
	GenChart       string    `gorm:"column:genChart;comment:生成的图表数据" json:"genChart"`                                       // 生成的图表数据
	GenResult      string    `gorm:"column:genResult;comment:生成的分析结论" json:"genResult"`                                     // 生成的分析结论
	Status         string    `gorm:"column:status;not null;default:wait;comment:wait,running,succeed,failed" json:"status"` // wait,running,succeed,failed
	ExecMessage    string    `gorm:"column:execMessage;comment:执行信息" json:"execMessage"`                                    // 执行信息
	UserID         int64     `gorm:"column:userId;comment:创建用户 id" json:"userId"`                                           // 创建用户 id
	CreateTime     time.Time `gorm:"column:createTime;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"createTime"`   // 创建时间
	UpdateTime     time.Time `gorm:"column:updateTime;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updateTime"`   // 更新时间
	IsDelete       int32     `gorm:"column:isDelete;not null;comment:是否删除" json:"isDelete"`                                 // 是否删除
}

// TableName Chart's table name
func (*Chart) TableName() string {
	return TableNameChart
}
