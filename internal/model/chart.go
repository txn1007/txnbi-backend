package model

import (
	"time"
)

const TableNameChart = "chart"

// Chart 图表信息表
type Chart struct {
	ID             int64     `gorm:"column:id;primaryKey;autoIncrement;comment:ID" json:"id"`                                                                // ID
	Goal           string    `gorm:"column:goal;type:text;comment:分析目标" json:"goal"`                                                                         // 分析目标
	Name           string    `gorm:"column:name;type:varchar(128);comment:图表名称" json:"name"`                                                                 // 图表名称
	ChartTableName string    `gorm:"column:chartTableName;type:varchar(64);comment:用户原始数据的表名" json:"chartTableName"`                                         // 用户原始数据的表名
	ChartType      string    `gorm:"column:chartType;type:varchar(32);comment:图表类型" json:"chartType"`                                                        // 图表类型
	GenChart       string    `gorm:"column:genChart;type:text;comment:生成的图表数据" json:"genChart"`                                                              // 生成的图表数据
	GenResult      string    `gorm:"column:genResult;type:text;comment:生成的分析结论" json:"genResult"`                                                            // 生成的分析结论
	Status         string    `gorm:"column:status;type:varchar(16);not null;default:wait;comment:状态;index:idx_status" json:"status"`                         // 状态：wait,running,succeed,failed
	ExecMessage    string    `gorm:"column:execMessage;type:text;comment:执行信息" json:"execMessage"`                                                           // 执行信息
	UserID         int64     `gorm:"column:userId;type:bigint unsigned;comment:创建用户ID;index:idx_userId" json:"userId"`                                       // 创建用户ID
	CreateTime     time.Time `gorm:"column:createTime;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:创建时间;index:idx_createTime" json:"createTime"` // 创建时间
	UpdateTime     time.Time `gorm:"column:updateTime;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updateTime"`                      // 更新时间
	IsDelete       int8      `gorm:"column:isDelete;type:tinyint;not null;default:0;comment:是否删除" json:"isDelete"`                                           // 是否删除
}

// TableName Chart's table name
func (*Chart) TableName() string {
	return TableNameChart
}
