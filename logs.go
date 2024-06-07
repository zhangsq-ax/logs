package logs

import (
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

// 决定日志级别的环境变量名称
var logLevelEnvName = "ENV"

// 环境变量值与日志级别的对应表，如果不存在则使用默认日志级别
var envLogLevels = map[string]zapcore.Level{
	"dev": zapcore.DebugLevel,
}

// 默认日志级别
var defaultLogLevel = zapcore.InfoLevel

var writeSyncer zapcore.WriteSyncer = zapcore.AddSync(os.Stdout)

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

func getLogger(encoderFun func(encoderConfig zapcore.EncoderConfig) zapcore.Encoder) *zap.Logger {
	//core := zapcore.NewCore(encoderFun(config), zapcore.AddSync(os.Stdout), GetLogLevel())
	core := zapcore.NewCore(encoderFun(config), writeSyncer, GetLogLevel())
	return zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel))
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
