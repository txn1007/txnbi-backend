package gorm_gen

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"txnbi-backend/conf"
)

// 请在项目根目录下运行
func GenModel() {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "./internal/gorm-gen",
		ModelPkgPath: "./internal/model",
	})

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s", conf.MySQLCfg.User, conf.MySQLCfg.Password,
		conf.MySQLCfg.Host, conf.MySQLCfg.Port, conf.MySQLCfg.DBName, conf.MySQLCfg.Charset, conf.MySQLCfg.ParseTime, conf.MySQLCfg.Location)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	g.UseDB(db)

	g.ApplyBasic(g.GenerateModel("chart"), g.GenerateModel("user"))
	g.Execute()
}
