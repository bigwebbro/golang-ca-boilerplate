package config

import "github.com/spf13/viper"

func Viper() *viper.Viper {
	v := viper.New()
	v.SetEnvPrefix("APP")
	v.AutomaticEnv()

	return v
}
