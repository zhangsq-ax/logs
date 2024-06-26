package logs

func Debug(args ...interface{}) {
	Logger().Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	Logger().Debugf(template, args...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	Logger().Debugw(msg, keysAndValues...)
}

func Info(args ...interface{}) {
	Logger().Info(args...)
}

func Infof(template string, args ...interface{}) {
	Logger().Infof(template, args...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	Logger().Infow(msg, keysAndValues...)
}

func Warn(args ...interface{}) {
	Logger().Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	Logger().Warnf(template, args...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	Logger().Warnw(msg, keysAndValues...)
}

func Error(args ...interface{}) {
	Logger().Error(args...)
}

func Errorf(template string, args ...interface{}) {
	Logger().Errorf(template, args...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	Logger().Errorw(msg, keysAndValues...)
}

func Fatal(args ...interface{}) {
	Logger().Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	Logger().Fatalf(template, args...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	Logger().Fatalw(msg, keysAndValues...)
}

func Panic(args ...interface{}) {
	Logger().Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	Logger().Panicf(template, args...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	Logger().Panicw(msg, keysAndValues...)
}

func ConsoleDebug(args ...interface{}) {
	ConsoleLogger().Debug(args...)
}

func ConsoleDebugf(template string, args ...interface{}) {
	ConsoleLogger().Debugf(template, args...)
}

func ConsoleDebugw(msg string, keysAndValues ...interface{}) {
	ConsoleLogger().Debugw(msg, keysAndValues...)
}

func ConsoleInfo(args ...interface{}) {
	ConsoleLogger().Info(args...)
}

func ConsoleInfof(template string, args ...interface{}) {
	ConsoleLogger().Infof(template, args...)
}

func ConsoleInfow(msg string, keysAndValues ...interface{}) {
	ConsoleLogger().Infow(msg, keysAndValues...)
}

func ConsoleWarn(args ...interface{}) {
	ConsoleLogger().Warn(args...)
}

func ConsoleWarnf(template string, args ...interface{}) {
	ConsoleLogger().Warnf(template, args...)
}

func ConsoleWarnw(msg string, keysAndValues ...interface{}) {
	ConsoleLogger().Warnw(msg, keysAndValues...)
}

func ConsoleError(args ...interface{}) {
	ConsoleLogger().Error(args...)
}

func ConsoleErrorf(template string, args ...interface{}) {
	ConsoleLogger().Errorf(template, args...)
}

func ConsoleErrorw(msg string, keysAndValues ...interface{}) {
	ConsoleLogger().Errorw(msg, keysAndValues...)
}

func ConsoleFatal(args ...interface{}) {
	ConsoleLogger().Fatal(args...)
}

func ConsoleFatalf(template string, args ...interface{}) {
	ConsoleLogger().Fatalf(template, args...)
}

func ConsoleFatalw(msg string, keysAndValues ...interface{}) {
	ConsoleLogger().Fatalw(msg, keysAndValues...)
}

func ConsolePanic(args ...interface{}) {
	ConsoleLogger().Panic(args...)
}

func ConsolePanicf(template string, args ...interface{}) {
	ConsoleLogger().Panicf(template, args...)
}

func ConsolePanicw(msg string, keysAndValues ...interface{}) {
	ConsoleLogger().Panicw(msg, keysAndValues...)
}
