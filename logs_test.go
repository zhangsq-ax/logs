package logs

import (
	"fmt"
	"go.uber.org/zap/zapcore"
	"os"
	"testing"
	"time"
)

/*func TestConfigEnvLogLevels(t *testing.T) {
	// 默认
	level := GetLogLevel()
	assert.Equal(t, zapcore.InfoLevel, level)

	// 默认 dev
	err := os.Setenv("ENV", "dev")
	assert.NoError(t, err)
	level = GetLogLevel()
	assert.Equal(t, zapcore.DebugLevel, level)

	// 自定义 prod
	SetEnvLogLevel("prod", zapcore.WarnLevel)
	err = os.Setenv("ENV", "prod")
	level = GetLogLevel()
	assert.Equal(t, zapcore.WarnLevel, level)

	// 恢复环境变量
	err = os.Unsetenv("ENV")
	assert.NoError(t, err)
}

func TestNewConsoleLogger(t *testing.T) {
	var logger = NewConsoleLogger()
	logger.Info("test")
	logger.Infow("test", zap.String("foo", "bar"))
}

func TestNewJSONLogger(t *testing.T) {
	var logger = NewJSONLogger("somename")
	logger.Infow("test", zap.String("foo", "bar"))
}

func TestConsoleErrorw(t *testing.T) {
	ConsoleErrorw("console-error", zap.String("foo", "bar"), zap.Int("number", 5))
}

func TestErrorw(t *testing.T) {
	Errorw("json-error", zap.String("foo", "bar"), zap.Int("number", 6))
}

func TestSetLogFile(t *testing.T) {
	SetLogFile("test.log", 1, 3, 10, true)
	logger := NewJSONLogger("somename")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			logger.Infow("test", zap.String("foo", "bar"))
		}
	}
}*/

func TestLogger(t *testing.T) {
	go func() {
		for {
			fmt.Println("--------------------------")
			Logger().Debug("this is debug")
			Logger().Info("this is info")
			Logger().Warn("this is warn")
			Logger().Error("this is error")
			Debug("this is debug")
			Info("this is info")
			Warn("this is warn")
			Error("this is error")
			ConsoleDebug("this is console debug")
			ConsoleInfo("this is console info")
			ConsoleWarn("this is console warn")
			ConsoleError("this is console error")
			fmt.Println("===========================")
			time.Sleep(5 * time.Second)
		}
	}()
}

func TestStartMonitorLogLevel(t *testing.T) {
	StartMonitorLogLevel(5 * time.Second)
	time.Sleep(10 * time.Second)
	t.Log("set dev")
	_ = os.Setenv("ENV", "dev")
	time.Sleep(10 * time.Second)
	t.Log("set test")
	_ = os.Setenv("ENV", "test")
	time.Sleep(10 * time.Second)
	t.Log("set prod")
	_ = os.Setenv("ENV", "prod")
	time.Sleep(10 * time.Second)
	t.Log("unset")
	_ = os.Unsetenv("ENV")
	time.Sleep(20 * time.Second)
}

func TestSetDefaultLogLevel(t *testing.T) {
	t.Log("set debug level")
	SetDefaultLogLevel(zapcore.DebugLevel)
	time.Sleep(10 * time.Second)
	t.Log("set info level")
	SetDefaultLogLevel(zapcore.InfoLevel)
	time.Sleep(10 * time.Second)
	t.Log("set warn level")
	SetDefaultLogLevel(zapcore.WarnLevel)
	time.Sleep(10 * time.Second)
	t.Log("set error level")
	SetDefaultLogLevel(zapcore.ErrorLevel)
	time.Sleep(20 * time.Second)
}
