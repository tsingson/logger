package logger

import (
	"time"

	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/spf13/afero"
	"go.uber.org/zap/zapcore"
)

var defaultLumberJack = &lumberjack.Logger{
	Filename:   "uber-zapadapter.log",
	MaxSize:    100, // megabytes
	MaxBackups: 31,
	MaxAge:     31,    // days
	Compress:   false, // 开发时不压缩
}

// const (
// 	cacheSize = 1024 * 1024 * 4
// 	timeOut   = time.Duration(50) * time.Millisecond
// )

// newZapCore initial a zapadapter log
func newZapCore(path, prefix string, level zapcore.LevelEnabler) zapcore.Core {
	dataTimeFmtInFileName := time.Now().Format("2006-01-02-15")

	var logPath string

	if len(path) == 0 {
		logPath = defaultLogPath(subPath)
	} else {
		logPath = path
	}

	afs := afero.NewOsFs()
	check, _ := afero.DirExists(afs, logPath)
	if !check {
		err := afs.MkdirAll(logPath, 0o755)
		if err != nil {
			panic("can't make path" + logPath)
		}
	}

	var logFilename string
	if len(prefix) == 0 {
		prefix = "default"
	}

	logFilename = logPath + "/" + prefix + "-" + dataTimeFmtInFileName + ".log"

	defaultLumberJack.Filename = logFilename

	w := &zapcore.BufferedWriteSyncer{
		WS:   zapcore.AddSync(defaultLumberJack),
		Size: _bufferSize,
	}

	return zapcore.NewCore(zapcore.NewJSONEncoder(defaultEncoder), w, level)
}

var defaultEncoder = zapcore.EncoderConfig{
	TimeKey:        "time",
	LevelKey:       "level",
	NameKey:        "name",
	CallerKey:      "caller",
	MessageKey:     "msg",
	StacktraceKey:  "stack",
	EncodeLevel:    zapcore.LowercaseLevelEncoder,
	EncodeTime:     zapcore.ISO8601TimeEncoder,
	EncodeDuration: zapcore.NanosDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
}
