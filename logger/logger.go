package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func init() {

	var err error

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""
	config := zap.NewProductionConfig()
	config.EncoderConfig = encoderConfig
	logger, err = config.Build(zap.AddCallerSkip(1))

	defer logger.Sync() // flushes buffer, if any
	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zapcore.Field) {
	logger.Info(message, fields...)
}

func Debug(message string, fields ...zapcore.Field) {
	logger.Debug(message, fields...)
}

func Error(message string, fields ...zapcore.Field) {
	logger.Error(message, fields...)
}
