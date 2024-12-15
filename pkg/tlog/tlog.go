package tlog

import "C"
import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

var L zerolog.Logger
var logFile *os.File

func init() {
	// 将日志写入本地文件和标准错误输出
	f, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	logFile = f
	r := zerolog.MultiLevelWriter(logFile, os.Stderr)

	zerolog.TimeFieldFormat = time.RFC3339
	L = zerolog.New(r).With().Timestamp().Caller().Logger()
	log.Logger = L

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
	return logFile.Close()
}
