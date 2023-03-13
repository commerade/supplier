package config

import (
	"github.com/spf13/viper"
)

func LoadEnvironmentVariables() {
	viper.AddConfigPath("./")
	viper.SetConfigName(APP)
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		viper.AutomaticEnv()
	}
}
