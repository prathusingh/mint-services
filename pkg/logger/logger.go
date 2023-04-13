package logger

import (
	"github.com/prathusingh/mint-services/pkg/constants"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

type Config struct {
	LogLevel string `mapstructure:"level"`
	DevMode bool `mapstructure:"devMode"`
	Encoder string `mapstructure:"encoder"`
}