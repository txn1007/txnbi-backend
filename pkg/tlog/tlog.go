package tlog

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var logger *lumberjack.Logger

func init() {
	logger = &lumberjack.Logger{
		Filename:   "app.log", // 日志文件名
		MaxSize:    100,       // 每个日志文件的最大大小（MB）
		MaxBackups: 5,         // 保留的最大备份数
		MaxAge:     7 * 5,     // 保留日志的最大天数（两周）
		Compress:   true,      // 是否压缩旧日志文件
	}
	r := zerolog.MultiLevelWriter(logger, os.Stderr)

	zerolog.TimeFieldFormat = time.RFC3339
	l := zerolog.New(r).With().Timestamp().Caller().Logger()
	log.Logger = l

	//// todo 优化输出结构
	//r := zerolog.ConsoleWriter{
	//	Out:        os.Stderr,
	//	TimeFormat: time.RFC3339,
	//	FormatLevel: func(i interface{}) string {
	//		return strings.ToUpper(fmt.Sprintf("[%s]", i))
	//	},
	//	FormatMessage: func(i interface{}) string {
	//		return fmt.Sprintf("%s", i)
	//	},
	//}

	//L = zerolog.New(r).With().Timestamp().Logger()
}

func CloseLogFile() error {
	return logger.Close()
}
