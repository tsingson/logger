package logger

import (
	"go.uber.org/zap/zapcore"
)

// SizeLogger size logger
func sizeLogger(path, prefix string, debug bool) zapcore.Core {
	if debug {
		// opts = append(opts, zapadapter.AddCaller())
		// opts = append(opts, zapadapter.AddStacktrace(zapadapter.WarnLevel))

		stdoutWs := &zapcore.BufferedWriteSyncer{
			WS:   zapcore.AddSync(getStdout()),
			Size: _bufferSize,
		}

		return zapcore.NewTee(
			zapcore.NewCore(zapcore.NewConsoleEncoder(defaultEncoder), stdoutWs, _defaultLevel),
			newZapCore(path, prefix, _defaultLevel))
	}
	return newZapCore(path, prefix, _defaultLevel)
}
