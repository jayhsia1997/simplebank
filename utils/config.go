package utils

import (
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURL string `mapstructure:"DATABASE_URL"`
	Host        string `mapstructure:"HOST"`
	Port        string `mapstructure:"PORT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigFile(filepath.Join(path, ".env"))
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
