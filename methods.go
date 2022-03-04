package logs

import "go.uber.org/zap"

var jsonLogger *zap.SugaredLogger
var consoleLogger *zap.SugaredLogger

func getJSONLogger() *zap.SugaredLogger {
	if jsonLogger == nil {
		jsonLogger = NewJSONLogger("__logs__")
	}
	return jsonLogger
}

func getConsoleLogger() *zap.SugaredLogger {
	if consoleLogger == nil {
		consoleLogger = NewConsoleLogger("__logs__")
	}
	return consoleLogger
}

func Debug(args ...interface{}) {
	logger := getJSONLogger()
	logger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	logger := getJSONLogger()
	logger.Debugf(template, args...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	logger := getJSONLogger()
	logger.Debugw(msg, keysAndValues...)
}

func Info(args ...interface{}) {
	logger := getJSONLogger()
	logger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	logger := getJSONLogger()
	logger.Infof(template, args...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	logger := getJSONLogger()
	logger.Infow(msg, keysAndValues...)
}

func Warn(args ...interface{}) {
	logger := getJSONLogger()
	logger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	logger := getJSONLogger()
	logger.Warnf(template, args...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	logger := getJSONLogger()
	logger.Warnw(msg, keysAndValues...)
}

func Error(args ...interface{}) {
	logger := getJSONLogger()
	logger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	logger := getJSONLogger()
	logger.Errorf(template, args...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	logger := getJSONLogger()
	logger.Errorw(msg, keysAndValues...)
}

func Fatal(args ...interface{}) {
	logger := getJSONLogger()
	logger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	logger := getJSONLogger()
	logger.Fatalf(template, args...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	logger := getJSONLogger()
	logger.Fatalw(msg, keysAndValues...)
}

func Panic(args ...interface{}) {
	logger := getJSONLogger()
	logger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	logger := getJSONLogger()
	logger.Panicf(template, args...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	logger := getJSONLogger()
	logger.Panicw(msg, keysAndValues...)
}

func ConsoleDebug(args ...interface{}) {
	logger := getConsoleLogger()
	logger.Debug(args...)
}

func ConsoleDebugf(template string, args ...interface{}) {
	logger := getConsoleLogger()
	logger.Debugf(template, args...)
}

func ConsoleDebugw(msg string, keysAndValues ...interface{}) {
	logger := getConsoleLogger()
	logger.Debugw(msg, keysAndValues...)
}

func ConsoleInfo(args ...interface{}) {
	logger := getConsoleLogger()
	logger.Info(args...)
}

func ConsoleInfof(template string, args ...interface{}) {
	logger := getConsoleLogger()
	logger.Infof(template, args...)
}

func ConsoleInfow(msg string, keysAndValues ...interface{}) {
	logger := getConsoleLogger()
	logger.Infow(msg, keysAndValues...)
}

func ConsoleWarn(args ...interface{}) {
	logger := getConsoleLogger()
	logger.Warn(args...)
}

func ConsoleWarnf(template string, args ...interface{}) {
	logger := getConsoleLogger()
	logger.Warnf(template, args...)
}

func ConsoleWarnw(msg string, keysAndValues ...interface{}) {
	logger := getConsoleLogger()
	logger.Warnw(msg, keysAndValues...)
}

func ConsoleError(args ...interface{}) {
	logger := getConsoleLogger()
	logger.Error(args...)
}

func ConsoleErrorf(template string, args ...interface{}) {
	logger := getConsoleLogger()
	logger.Errorf(template, args...)
}

func ConsoleErrorw(msg string, keysAndValues ...interface{}) {
	logger := getConsoleLogger()
	logger.Errorw(msg, keysAndValues...)
}

func ConsoleFatal(args ...interface{}) {
	logger := getConsoleLogger()
	logger.Fatal(args...)
}

func ConsoleFatalf(template string, args ...interface{}) {
	logger := getConsoleLogger()
	logger.Fatalf(template, args...)
}

func ConsoleFatalw(msg string, keysAndValues ...interface{}) {
	logger := getConsoleLogger()
	logger.Fatalw(msg, keysAndValues...)
}

func ConsolePanic(args ...interface{}) {
	logger := getConsoleLogger()
	logger.Panic(args...)
}

func ConsolePanicf(template string, args ...interface{}) {
	logger := getConsoleLogger()
	logger.Panicf(template, args...)
}

func ConsolePanicw(msg string, keysAndValues ...interface{}) {
	logger := getConsoleLogger()
	logger.Panicw(msg, keysAndValues...)
}
