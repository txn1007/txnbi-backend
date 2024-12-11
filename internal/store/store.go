package store

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"txnbi-backend/conf"
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
	
}
