package logs

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"testing"
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
