package logger

import (
	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func New(logLevel string) *zap.SugaredLogger {
	encoderConfig := ecszap.NewDefaultEncoderConfig()
	var core zapcore.Core
	switch logLevel {
	case "DEVELOPMENT":
		core = ecszap.NewCore(encoderConfig, os.Stdout, zap.DebugLevel)
	default:
		core = ecszap.NewCore(encoderConfig, os.Stdout, zap.InfoLevel)
	}
	logger := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logger)
	return logger.Sugar()
}
