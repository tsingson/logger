package logger

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog/diode"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/tsingson/logger/rotatelogs"
)

// DayLogger day logger
func dayLogger(path, prefix string, days int64, debug bool) zapcore.Core {
	// 实现两个判断日志等级的interface
	// 仅打印Error级别以上的日志
	_errorLevelPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	// // 打印所有级别的日志
	// lowPriority := zapadapter.LevelEnablerFunc(func(lvl zapcore.Level) bool {
	// 	return lvl >= zapcore.DebugLevel
	// })

	if len(path) == 0 {
		path = defaultLogPath(subPath)
	}
	if len(prefix) == 0 {
		prefix = os.Args[0]
	}
	// 获取 info、warn日志文件的io.Writer 抽象 getWriter() 在下方实现

	var core zapcore.Core
	if debug {
		allWriter := getWriter(path+"/"+prefix+".log", days)
		errorWriter := getWriter(path+"/"+prefix+"-error.log", days)

		stdoutWs := zapcore.AddSync(getStdout())

		allWs := &zapcore.BufferedWriteSyncer{
			WS:   zapcore.AddSync(allWriter),
			Size: _bufferSizedebug,
		}
		errorWs := &zapcore.BufferedWriteSyncer{
			WS:   zapcore.AddSync(errorWriter),
			Size: _bufferSizedebug,
		}
		// 最后创建具体的Logger
		core = zapcore.NewTee(
			zapcore.NewCore(zapcore.NewConsoleEncoder(defaultEncoder), stdoutWs, _defaultLevel),
			zapcore.NewCore(zapcore.NewJSONEncoder(defaultEncoder), allWs, _defaultLevel),
			zapcore.NewCore(zapcore.NewJSONEncoder(defaultEncoder), errorWs, _errorLevelPriority),
		)
	} else {
		allWriter := getWriter(path+"/"+prefix+".log", days)
		errorWriter := getWriter(path+"/"+prefix+"-error.log", days)

		allWs := &zapcore.BufferedWriteSyncer{
			WS:   zapcore.AddSync(allWriter),
			Size: _bufferSize,
		}
		errorWs := &zapcore.BufferedWriteSyncer{
			WS:   zapcore.AddSync(errorWriter),
			Size: _bufferSize,
		}
		core = zapcore.NewTee(
			zapcore.NewCore(zapcore.NewJSONEncoder(defaultEncoder), allWs, _defaultLevel),
			zapcore.NewCore(zapcore.NewJSONEncoder(defaultEncoder), errorWs, _errorLevelPriority),
		)
	}

	return core // , zapadapter.AddCaller()) // 需要传入 zapadapter.AddCaller() 才会显示打日志点的文件名和行数, 有点小坑
}

var _rotateLog *rotatelogs.RotateLogs

func getWriter(filename string, days int64) io.Writer {
	//
	// w := &zapcore.BufferedWriteSyncer{
	// WS: zapcore. AddSync(file),
	// }
	// 生成rotatelogs的Logger 实际生成的文件名 demo.log.YYmmddHH
	// demo.log是指向最新日志的链接
	// 保存7天内的日志，每1小时(整点)分割一次日志
	var err error
	_rotateLog, err = rotatelogs.New(
		filename+"-%Y%m%d%H"+".json", // 没有使用go风格反人类的format格式
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*time.Duration(24)*time.Duration(days)),
		rotatelogs.WithRotationTime(_intervalWrite),
	)
	if err != nil {
		return diode.NewWriter(os.Stdout, _bufferSize, _interval, func(missed int) {
			// fmt.Printf("Dropped %d messages\n", missed)
		})
	}

	return diode.NewWriter(_rotateLog, _bufferSize, _interval, func(missed int) {
		// fmt.Printf("Dropped %d messages\n", missed)
	})
}

func getStdout() io.Writer {
	w := diode.NewWriter(os.Stdout, _bufferSize, _interval, func(missed int) {
		// fmt.Printf("Dropped %d messages\n", missed)
	})
	return w
}
