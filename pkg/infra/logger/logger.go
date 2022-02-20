package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"jwemanager/pkg/app/interfaces"
)

func NewLogger() interfaces.ILogger {
	goEnv := os.Getenv("GO_ENV")

	zapLogLevel := getLogLevel()

	var zapInstance *zap.Logger
	switch goEnv {
	case "production":
	case "staging":
		zapInstance, _ = zap.NewProduction(zap.IncreaseLevel(zapLogLevel), zap.AddStacktrace(zap.ErrorLevel))
	case "development":
	case "test":
		config := zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		config.Level.Enabled(zapLogLevel)
		zapInstance, _ = config.Build()
	default:
		break
	}

	return zapInstance
}

func getLogLevel() zapcore.Level {
	logLevel := os.Getenv("LOG_LEVEL")
	switch logLevel {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "panic":
		return zap.PanicLevel
	default:
		return zap.InfoLevel
	}
}
