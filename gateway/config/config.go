package config

import (
	"flag"
	"fmt"
	"os"
	"github.com/prathusingh/mint-services/pkg/constants"

	"github.com/spf13/viper"
	"github.com/pkg/errors"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "gateway microservice config path")
}

type Config struct {
	ServiceName string `mapstructure:"serviceName`
}

func InitConfig() (*Config, error) {
	if configPath == "" {
		conigPathFromEnv := os.Getenv(constants.ConfigPath)
		if configPath != "" {
			configPath = conigPathFromEnv
		} else {
			getwd, err := os.Getwd()
			if err != nil {
				return nil, errors.Wrap(err, "os.Getwd")
			}
			configPath = fmt.Sprintf("%s/gateway/config/config.yml", getwd)
		}
	}

	cfg := &Config{}
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "viper.ReadInConfig")
	}

	if err := viper.Unmarshal(cfg);  err != nil {
		return nil, errors.Wrap(err, "viper.Unmarshal")
	}

	return cfg, nil
}
