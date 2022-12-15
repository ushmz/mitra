package config

import "github.com/spf13/viper"

var config *viper.Viper

// Init read config file and set variables
func Init(configName string) error {
	config = viper.New()
	config.SetConfigFile("yaml")
	config.SetConfigName(configName)
	config.AddConfigPath(".")
	config.AddConfigPath("config/")
	config.SetDefault("version", "beta")
	config.SetDefault("config", configName)
	if err := config.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

// GetConfig return viper configuration registry
func GetConfig() *viper.Viper {
	return config
}
