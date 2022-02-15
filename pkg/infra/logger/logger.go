package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger() *zap.Logger {
	goEnv := os.Getenv("GO_ENV")

	zapLogLevel := getLogLevel()

	var zapInstance *zap.Logger
	switch goEnv {
	case "production":
		zapInstance, _ = zap.NewProduction(zap.IncreaseLevel(zapLogLevel), zap.AddStacktrace(zap.ErrorLevel))
	case "staging":
		zapInstance, _ = zap.NewProduction(zap.IncreaseLevel(zapLogLevel), zap.AddStacktrace(zap.ErrorLevel))
	case "development":
		zapInstance, _ = zap.NewDevelopment(zap.IncreaseLevel(zapLogLevel), zap.AddStacktrace(zap.ErrorLevel))
	case "test":
		zapInstance, _ = zap.NewDevelopment(zap.IncreaseLevel(zapLogLevel), zap.AddStacktrace(zap.ErrorLevel))
	default:
		break
	}

	return zapInstance
}

func getLogLevel() zapcore.Level {
	logLevel := os.Getenv("LOG_LEVEL")
	switch logLevel {
	case "Debug":
		return zap.DebugLevel
	case "Info":
		return zap.InfoLevel
	case "Warn":
		return zap.WarnLevel
	case "Error":
		return zap.ErrorLevel
	case "Panic":
		return zap.PanicLevel
	default:
		return zap.InfoLevel
	}
}
