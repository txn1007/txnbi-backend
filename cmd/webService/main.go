package main

import (
	"fmt"
	"txnbi-backend/conf"
	_ "txnbi-backend/docs"
	txnbi "txnbi-backend/internal"
)

//	@title			txnbi API
//	@version		1.0
//	@description	txnbi 的 API
//	@host			localhost:8080
//	@BasePath		/

func main() {
	r := txnbi.Route()
	addr := fmt.Sprintf("%s:%d", conf.TxnBICfg.Host, conf.TxnBICfg.Port)
	err := r.Run(addr)
	if err != nil {
		panic(err)
	}
}

// 根据数据库已有的表结构生成 model
//func main() {
//	gorm_gen.GenModel()
//}
