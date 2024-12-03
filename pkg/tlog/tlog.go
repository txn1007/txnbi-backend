package tlog

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"strings"
	"time"
)

var L zerolog.Logger

func init() {
	r := zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
		FormatLevel: func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("[%s]", i))
		},
		FormatMessage: func(i interface{}) string {
			return fmt.Sprintf("%s", i)
		},
	}

	L = zerolog.New(r).With().Timestamp().Logger()
}
