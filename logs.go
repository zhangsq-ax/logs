package logs

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var config = zapcore.EncoderConfig{
	NameKey:      "module",
	MessageKey:   "msg",                       //结构化（json）输出：msg的key
	LevelKey:     "level",                     //结构化（json）输出：日志级别的key（INFO，WARN，ERROR等）
	TimeKey:      "ts",                        //结构化（json）输出：时间的key（INFO，WARN，ERROR等）
	CallerKey:    "file",                      //结构化（json）输出：打印日志的文件对应的Key
	EncodeLevel:  zapcore.CapitalLevelEncoder, //将日志级别转换成大写（INFO，WARN，ERROR等）
	EncodeCaller: zapcore.ShortCallerEncoder,  //采用短文件路径编码输出（test/main.go:14 ）
	EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}, //输出的时间格式
	EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendInt64(int64(d) / 1000000)
	}, //
}

var (
	jsonLoggers    = make(map[string]*zap.SugaredLogger)
	consoleLoggers = make(map[string]*zap.SugaredLogger)
)

// 决定日志级别的环境变量名称
var logLevelEnvName = "ENV"

// 环境变量值与日志级别的对应表，如果不存在则使用默认日志级别
var envLogLevels = map[string]zapcore.Level{
	"dev":  zapcore.DebugLevel,
	"test": zapcore.InfoLevel,
	"prod": zapcore.WarnLevel,
}

// 默认日志级别
var defaultLogLevel = zapcore.InfoLevel

var writeSyncer zapcore.WriteSyncer = zapcore.AddSync(os.Stdout)

var (
	monitorTicker   *time.Ticker
	monitorStopChan chan struct{}
)

func reset() {
	jsonLoggers = make(map[string]*zap.SugaredLogger)
	consoleLoggers = make(map[string]*zap.SugaredLogger)
}

// GetLogLevelEnvName 获取决定日志级别的环境变量名称
func GetLogLevelEnvName() string {
	return logLevelEnvName
}

// SetLogLevelEnvName 设置决定日志级别的环境变量名称
func SetLogLevelEnvName(envName string) {
	logLevelEnvName = envName
}

// GetEnvLogLevels 获取环境变量值与日志级别的对应表
func GetEnvLogLevels() map[string]zapcore.Level {
	return envLogLevels
}

// SetEnvLogLevels 设置环境变量值与日志级别的对应表
func SetEnvLogLevels(levels map[string]zapcore.Level) {
	envLogLevels = levels
}

// SetEnvLogLevel 设置单个环境变量值与日志级别的对应
func SetEnvLogLevel(envValue string, level zapcore.Level) {
	envLogLevels[envValue] = level
}

func SetDefaultLogLevel(level zapcore.Level) {
	reset()
	defaultLogLevel = level
}

func StartMonitorLogLevel(duration time.Duration) {
	if duration == 0 {
		duration = 10 * time.Second
	}
	go func() {
		monitorTicker = time.NewTicker(duration)
		monitorStopChan = make(chan struct{})
		defer close(monitorStopChan)
		for {
			select {
			case <-monitorTicker.C:
				level := GetLogLevel()
				if level != defaultLogLevel {
					reset()
				}
			case <-monitorStopChan:
				return
			}
		}
	}()
}

func StopMonitorLogLevel() {
	if monitorTicker != nil && monitorStopChan != nil {
		monitorTicker.Stop()
		close(monitorStopChan)
		monitorTicker = nil
		monitorStopChan = nil
	}
}

// GetLogLevel 获取当前的日志级别
func GetLogLevel() zapcore.Level {
	if level, ok := envLogLevels[os.Getenv(logLevelEnvName)]; ok {
		return level
	} else {
		return defaultLogLevel
	}
}

func SetWriteSyncer(w *zapcore.WriteSyncer) {
	if w == nil {
		writeSyncer = zapcore.AddSync(os.Stdout)
	} else {
		writeSyncer = *w
	}
}

// SetLogFile 设置日志文件，设置完成后日志会输出到此文件
// logFile: 日志文件路径
// maxSize：日志文件的最大大小，单位: MB
// maxBackups：旧日志文件保留的最大个数
// maxAge：旧日志文件保留的最大天数
// compress：是否压缩旧日志文件，默认为 false
func SetLogFile(logFile string, maxSize int, maxBackups int, maxAge int, compress bool) {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   compress,
	})
	SetWriteSyncer(&w)
}

func getLogger(encoderFun func(encoderConfig zapcore.EncoderConfig) zapcore.Encoder, callerSkips ...int) *zap.Logger {
	core := zapcore.NewCore(encoderFun(config), writeSyncer, GetLogLevel())
	callerSkip := 0
	if len(callerSkips) > 0 {
		callerSkip = callerSkips[0]
	}
	return zap.New(
		core,
		zap.AddCaller(),
		zap.AddStacktrace(zap.WarnLevel),
		zap.AddCallerSkip(callerSkip),
	)
}

// NewConsoleLogger 创建控制台格式日志器
func NewConsoleLogger(names ...string) *zap.SugaredLogger {
	logger := getLogger(zapcore.NewConsoleEncoder)
	if len(names) > 0 {
		name := names[0]
		if name != "" {
			logger.Named(name)
		}
	}
	return logger.Sugar()
}

func NewConsoleLoggerWithCallerSkip(skip int, names ...string) *zap.SugaredLogger {
	logger := getLogger(zapcore.NewConsoleEncoder, skip)
	if len(names) > 0 {
		name := names[0]
		if name != "" {
			logger.Named(name)
		}
	}
	return logger.Sugar()
}

// NewJSONLogger 创建 JSON 格式日志器
func NewJSONLogger(names ...string) *zap.SugaredLogger {
	logger := getLogger(zapcore.NewJSONEncoder)
	if len(names) > 0 {
		name := names[0]
		if name != "" {
			logger.Named(name)
		}
	}
	return logger.Sugar()
}

func NewJSONLoggerWithCallerSkip(skip int, names ...string) *zap.SugaredLogger {
	logger := getLogger(zapcore.NewJSONEncoder, skip)
	if len(names) > 0 {
		name := names[0]
		if name != "" {
			logger.Named(name)
		}
	}
	return logger.Sugar()
}

func JSONLogger(names ...string) *zap.SugaredLogger {
	var (
		logger *zap.SugaredLogger
		ok     bool
		name   string
	)
	if len(names) > 0 {
		name = names[0]
	} else {
		name = "default"
	}
	if logger, ok = jsonLoggers[name]; !ok {
		logger = NewJSONLogger(name)
		jsonLoggers[name] = logger
	}
	return logger
}

func JSONLoggerWithCallerSkip(skip int, names ...string) *zap.SugaredLogger {
	var (
		logger *zap.SugaredLogger
		ok     bool
		name   string
	)
	if len(names) > 0 {
		name = names[0]
	} else {
		name = fmt.Sprintf("skip_%d", skip)
	}
	if logger, ok = jsonLoggers[name]; !ok {
		logger = NewJSONLoggerWithCallerSkip(skip, name)
		jsonLoggers[name] = logger
	}
	return logger
}

func ConsoleLogger(names ...string) *zap.SugaredLogger {
	var (
		logger *zap.SugaredLogger
		ok     bool
		name   string
	)
	if len(names) > 0 {
		name = names[0]
	} else {
		name = "default"
	}
	if logger, ok = consoleLoggers[name]; !ok {
		logger = NewConsoleLogger(name)
		consoleLoggers[name] = logger
	}
	return logger
}

func ConsoleLoggerWithCallerSkip(skip int, names ...string) *zap.SugaredLogger {
	var (
		logger *zap.SugaredLogger
		ok     bool
		name   string
	)
	if len(names) > 0 {
		name = names[0]
	} else {
		name = fmt.Sprintf("skip_%d", skip)
	}
	if logger, ok = consoleLoggers[name]; !ok {
		logger = NewConsoleLoggerWithCallerSkip(skip, name)
		consoleLoggers[name] = logger
	}
	return logger
}

func Logger(names ...string) *zap.SugaredLogger {
	return JSONLogger(names...)
}

func CallerSkipLogger(skip int, names ...string) *zap.SugaredLogger {
	return JSONLoggerWithCallerSkip(skip, names...)
}
