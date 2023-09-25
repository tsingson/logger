package logger

import (
	"time"
)

const (
	subPath          = "log"
	_interval        = time.Duration(15) * time.Microsecond
	_intervalWrite   = time.Duration(5) * time.Minute
	_bufferSize      = 1024 * 1024 * 128
	_bufferSizedebug = 1024
)
