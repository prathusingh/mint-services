package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	LogLevel string `mapstructure:"level"`
	DevMode  bool   `mapstructure:"devMode"`
	Encoder  string `mapstructure:"encoder"`
}

func NewLoggerConfig(logLevel string, devMode bool, encoder string) *Config {
	return &Config{LogLevel: logLevel, DevMode: devMode, Encoder: encoder}
}

type Logger interface {
	InitLogger()
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
}

type appLogger struct {
	level       string
	devMode     bool
	encoding    string
	sugarLogger *zap.SugaredLogger
	logger      *zap.Logger
}

func NewAppLogger(cfg *Config) *appLogger {
	return &appLogger{level: cfg.LogLevel, devMode: cfg.DevMode, encoding: cfg.Encoder}
}

var loggerLevelMap = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"fatal": zapcore.FatalLevel,
}

func (l *appLogger) getLoggerLevel() zapcore.Level {
	level, exist := loggerLevelMap[l.level]
	if !exist {
		return zapcore.DebugLevel
	}

	return level
}

func (l *appLogger) InitLogger() {
	logLevel := l.getLoggerLevel()

	logWriter := zapcore.AddSync(os.Stdout)

	var encoderConfig zapcore.EncoderConfig

	if l.devMode {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderConfig = zap.NewProductionEncoderConfig()
	}

	var encoder zapcore.Encoder

	if l.encoding == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	core := zapcore.NewCore(encoder, logWriter, zap.NewAtomicLevelAt(logLevel))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	l.logger = logger
	l.sugarLogger = logger.Sugar()
}

func (l *appLogger) Debug(args ...interface{}) {
	l.sugarLogger.Debug(args...)
}

func (l *appLogger) Info(args ...interface{}) {
	l.sugarLogger.Info(args...)
}

func (l *appLogger) Warn(args ...interface{}) {
	l.sugarLogger.Warn(args...)
}

func (l *appLogger) Error(args ...interface{}) {
	l.sugarLogger.Error(args...)
}

func (l *appLogger) Fatal(args ...interface{}) {
	l.sugarLogger.Fatal(args...)
}
