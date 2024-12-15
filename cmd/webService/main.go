package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"txnbi-backend/conf"
	_ "txnbi-backend/docs"
	txnbi "txnbi-backend/internal"
	"txnbi-backend/pkg/tlog"
)

//	@title			txnbi API
//	@version		1.0
//	@description	txnbi 的 API
//	@host			localhost:8080
//	@BasePath		/

func main() {
	// 启动服务器
	// 创建路由并启动服务器
	r := txnbi.Route()
	addr := fmt.Sprintf("%s:%d", conf.TxnBICfg.Host, conf.TxnBICfg.Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}
	go func() {
		log.Info().Msgf("服务器启动，监听端口: %d", conf.TxnBICfg.Port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("服务器启动失败")
		}
	}()
	// 关闭服务器
	// 监听信号，当收到信号后进行关停
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// 创建一个带超时的上下文，尝试优雅地关闭服务器
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("优雅关停失败")
	}
	log.Info().Msg("关闭服务器成功")

	// 关闭日志文件句柄
	if err := tlog.CloseLogFile(); err != nil {
		panic(err)
	}

}

// 根据数据库已有的表结构生成 model
//func main() {
//	gorm_gen.GenModel()
//}
