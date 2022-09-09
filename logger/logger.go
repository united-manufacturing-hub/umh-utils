package logger

import (
	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// New creates a new logger with the given log level.
// envKey is the name of the environment variable that contains the log level.
func New(envKey string) *zap.SugaredLogger {
	var logLevel = os.Getenv(envKey)
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
