package logs

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"testing"
	"time"
)

func TestConfigEnvLogLevels(t *testing.T) {
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
}
