package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger() (*zap.Logger, error) {

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""
	config := zap.NewProductionConfig()
	config.EncoderConfig = encoderConfig
	return config.Build(zap.AddCallerSkip(1))

	// defer logger.Sync() // flushes buffer, if any
}
