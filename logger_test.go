package logger

import (
	"testing"

	"go.uber.org/zap/zapcore"
)

func TestNew(t *testing.T) {
	log := New(WithDebug(), WithLevel(zapcore.DebugLevel), WithPath("testdata/log"), WithPrefix("test"))
	log.Debug("test")
	defer func() {
		log.Sync()
	}()
}
