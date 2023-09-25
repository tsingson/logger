package logger

// MyLogger interface
type MyLogger interface {
	Fatalf(string, ...interface{})
	Debugf(string, ...interface{})
	Errorf(string, ...interface{})
	Infof(string, ...interface{})
	Warnf(string, ...interface{})
	Debug(...interface{})
	Warn(...interface{})
	Info(...interface{})
	Fatal(...interface{})
}

// PrintfLogger for fasthttp log interface
type PrintfLogger interface {
	Printf(format string, args ...interface{})
}
