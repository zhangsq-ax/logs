package logs

func Debug(args ...interface{}) {
	CallerSkipLogger(1).Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	CallerSkipLogger(1).Debugf(template, args...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	CallerSkipLogger(1).Debugw(msg, keysAndValues...)
}

func Info(args ...interface{}) {
	CallerSkipLogger(1).Info(args...)
}

func Infof(template string, args ...interface{}) {
	CallerSkipLogger(1).Infof(template, args...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	CallerSkipLogger(1).Infow(msg, keysAndValues...)
}

func Warn(args ...interface{}) {
	CallerSkipLogger(1).Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	CallerSkipLogger(1).Warnf(template, args...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	CallerSkipLogger(1).Warnw(msg, keysAndValues...)
}

func Error(args ...interface{}) {
	CallerSkipLogger(1).Error(args...)
}

func Errorf(template string, args ...interface{}) {
	CallerSkipLogger(1).Errorf(template, args...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	CallerSkipLogger(1).Errorw(msg, keysAndValues...)
}

func Fatal(args ...interface{}) {
	CallerSkipLogger(1).Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	CallerSkipLogger(1).Fatalf(template, args...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	CallerSkipLogger(1).Fatalw(msg, keysAndValues...)
}

func Panic(args ...interface{}) {
	CallerSkipLogger(1).Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	CallerSkipLogger(1).Panicf(template, args...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	CallerSkipLogger(1).Panicw(msg, keysAndValues...)
}

func ConsoleDebug(args ...interface{}) {
	ConsoleLoggerWithCallerSkip(1).Debug(args...)
}

func ConsoleDebugf(template string, args ...interface{}) {
	ConsoleLoggerWithCallerSkip(1).Debugf(template, args...)
}

func ConsoleDebugw(msg string, keysAndValues ...interface{}) {
	ConsoleLoggerWithCallerSkip(1).Debugw(msg, keysAndValues...)
}

func ConsoleInfo(args ...interface{}) {
	ConsoleLoggerWithCallerSkip(1).Info(args...)
}

func ConsoleInfof(template string, args ...interface{}) {
	ConsoleLoggerWithCallerSkip(1).Infof(template, args...)
}

func ConsoleInfow(msg string, keysAndValues ...interface{}) {
	ConsoleLoggerWithCallerSkip(1).Infow(msg, keysAndValues...)
}

func ConsoleWarn(args ...interface{}) {
	ConsoleLoggerWithCallerSkip(1).Warn(args...)
}

func ConsoleWarnf(template string, args ...interface{}) {
	ConsoleLoggerWithCallerSkip(1).Warnf(template, args...)
}

func ConsoleWarnw(msg string, keysAndValues ...interface{}) {
	ConsoleLoggerWithCallerSkip(1).Warnw(msg, keysAndValues...)
}

func ConsoleError(args ...interface{}) {
	ConsoleLoggerWithCallerSkip(1).Error(args...)
}

func ConsoleErrorf(template string, args ...interface{}) {
	ConsoleLoggerWithCallerSkip(1).Errorf(template, args...)
}

func ConsoleErrorw(msg string, keysAndValues ...interface{}) {
	ConsoleLoggerWithCallerSkip(1).Errorw(msg, keysAndValues...)
}

func ConsoleFatal(args ...interface{}) {
	ConsoleLoggerWithCallerSkip(1).Fatal(args...)
}

func ConsoleFatalf(template string, args ...interface{}) {
	ConsoleLoggerWithCallerSkip(1).Fatalf(template, args...)
}

func ConsoleFatalw(msg string, keysAndValues ...interface{}) {
	ConsoleLoggerWithCallerSkip(1).Fatalw(msg, keysAndValues...)
}

func ConsolePanic(args ...interface{}) {
	ConsoleLoggerWithCallerSkip(1).Panic(args...)
}

func ConsolePanicf(template string, args ...interface{}) {
	ConsoleLoggerWithCallerSkip(1).Panicf(template, args...)
}

func ConsolePanicw(msg string, keysAndValues ...interface{}) {
	ConsoleLoggerWithCallerSkip(1).Panicw(msg, keysAndValues...)
}
