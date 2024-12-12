package store

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"txnbi-backend/conf"
	"txnbi-backend/internal/model"
	"txnbi-backend/internal/store/myRedis"
)

var (
	DB *gorm.DB
)

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s",
		conf.MySQLCfg.User, conf.MySQLCfg.Password, conf.MySQLCfg.Host, conf.MySQLCfg.Port,
		conf.MySQLCfg.DBName, conf.MySQLCfg.Charset, conf.MySQLCfg.ParseTime, conf.MySQLCfg.Location)

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		fmt.Println("初始化全局DB失败！")
		panic(err)
	}
	DB = db
	fmt.Println("初始化全局DB成功！")

	// 缓存示例图表
	ExampleChartInDB := make([]model.ChartExample, 0)
	// 从 DB 中获取数据
	err = DB.Find(&ExampleChartInDB).Error
	if err != nil {
		panic(fmt.Sprintf("从 DB 层中查询示例图表数据失败！%s", err.Error()))
		return
	}
	// 将数据存储至 Redis
	err = myRedis.SetExampleChart(context.Background(), ExampleChartInDB)
	if err != nil {
		panic(fmt.Sprintf("将示例图表数据缓存到 Redis 失败！%s", err.Error()))
		return
	}

}
