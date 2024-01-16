package logger

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger zap adapter log
type Logger struct {
	core         zapcore.Core
	path         string
	prefix       string
	Log          *zap.Logger
	debug        bool
	storeInDay   bool
	addCaller    bool
	samplingCore bool

	logLevel zapcore.LevelEnabler
	days     int64
}

// ZapLogger logger
type ZapLogger = Logger

// An Option configures a Logger.
type Option interface {
	apply(*Logger)
}

// optionFunc wraps a func so that it satisfies the Option interface.
type optionFunc func(*Logger)

func (f optionFunc) apply(log *Logger) {
	f(log)
}

func (s *Logger) clone() *Logger {
	c := *s
	return &c
}

// WithOptions clones the current Logger, applies the supplied Options, and
// returns the resulting Logger. It's safe to use concurrently.
func (s *Logger) WithOptions(opts ...Option) *Logger {
	c := s.clone()
	for _, opt := range opts {
		opt.apply(c)
	}
	return c
}

// ZapLoggerOption options
type ZapLoggerOption func(*ZapLogger)

// WithDebug set debug
func WithDebug() ZapLoggerOption {
	return func(o *ZapLogger) {
		o.debug = true
	}
}

// WithDays set debug
func WithDays(days int64) ZapLoggerOption {
	return func(o *ZapLogger) {
		o.storeInDay = true
		o.days = days
	}
}

func WithSamplingCore() ZapLoggerOption {
	return func(logger *ZapLogger) {
		logger.samplingCore = true
	}
}

// WithStoreInDay set debug
func WithStoreInDay() ZapLoggerOption {
	return func(o *ZapLogger) {
		o.storeInDay = true
	}
}

// WithCaller set debug
func WithCaller() ZapLoggerOption {
	return func(o *ZapLogger) {
		o.addCaller = true
	}
}

// WithLevel set log level
func WithLevel(level zapcore.Level) ZapLoggerOption {
	return func(o *ZapLogger) {
		_defaultLevel.SetLevel(level)
		o.logLevel = _defaultLevel
	}
}

// WithPath set log level
func WithPath(path string) ZapLoggerOption {
	return func(o *ZapLogger) {
		o.path = path
	}
}

// WithPrefix set log level
func WithPrefix(prefix string) ZapLoggerOption {
	return func(o *ZapLogger) {
		o.prefix = prefix
	}
}

var defaultZapLogger = &ZapLogger{
	debug:      false,
	addCaller:  false,
	days:       int64(31),
	storeInDay: true,
	logLevel:   _defaultLevel,
}

// New init a log
func New(opts ...ZapLoggerOption) *ZapLogger {
	s := defaultZapLogger

	for _, o := range opts {
		o(s)
	}

	if len(s.path) == 0 {
		s.path = defaultLogPath(subPath)
	}

	if len(s.prefix) == 0 {
		s.prefix = StrBuilder(os.Args[0], "-", fmt.Sprint(os.Getpid()))
	}

	var core zapcore.Core
	if s.storeInDay {
		core = dayLogger(s.path, s.prefix, s.days, s.debug)
	} else {
		core = sizeLogger(s.path, s.prefix, s.debug)
	}
	s.core = core

	if s.addCaller {
		s.Log = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel))
	} else {
		if s.samplingCore {
			samplingCore := zapcore.NewSamplerWithOptions(
				core,
				time.Second, // interval
				3,           // log first 3 entries
				0,           // thereafter log zero entires within the interval
			)
			s.core = samplingCore
			s.Log = zap.New(samplingCore)
			return s
		} else {
			s.core = core
			s.Log = zap.New(core)
		}
	}

	return s
}

func (s *Logger) Core() zapcore.Core {
	return s.core
}

// Sync wrap sync
func (s *Logger) Sync() {
	_ = s.Log.Sync()
	if _rotateLog != nil {
		_ = _rotateLog.Rotate()
	}
	_ = s.Log.Sync()
}

// Named name logger
func (s *Logger) Named(str string) *zap.Logger {
	return s.Log.Named(str)
}
