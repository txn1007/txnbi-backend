package tlog

import "testing"

func TestLogPrintInFile(t *testing.T) {
	for i := 0; i < 100; i++ {
		L.Debug().Msg("123")
	}
}
