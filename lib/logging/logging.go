package logging

import (
	"log"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Level string
}

type Logger struct {
	*zap.SugaredLogger
}

const (
	defaultLevel = "info"
)

var config *Config
var sugar *Logger

var (
	defaultZapLogLevel = zapcore.InfoLevel
)

func New() *Logger {
	v := viper.New()
	v.AutomaticEnv()
	v.SetDefault("LOG_LEVEL", defaultLevel)

	config = &Config{
		Level: v.GetString("LOG_LEVEL"),
	}

	return InitLogger()
}

func InitLogger() *Logger {
	if sugar != nil {
		return sugar
	}
	// Initialize logger
	zapConfig := zap.NewProductionConfig()

	logLevel, err := zapcore.ParseLevel(config.Level)
	if err != nil {
		logLevel = defaultZapLogLevel
	}

	zapConfig.Level = zap.NewAtomicLevelAt(logLevel)
	zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, err := zapConfig.Build()
	if err != nil {
		log.Fatal("Failed to initiate zap logger! err:", err)
	}

	sugar = &Logger{
		logger.Sugar(),
	}

	logger.Debug("Successfully initialized Zap Logger!")
	return sugar
}

func SugarLog() *Logger {
	return sugar
}
