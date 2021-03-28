package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	PostgresUsername string `mapstructure:"POSTGRES_USERNAME"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDatabase string `mapstructure:"POSTGRES_DATABASE"`
	PostgresServerAddress string `mapstructure:"POSTGRES_SERVER_ADDRESS"`
	PostgresServerPort string `mapstructure:"POSTGRES_SERVER_PORT"`
	PostgresSslMode string `mapstructure:"POSTGRES_SSL_MODE"`
	Port string `mapstructure:"PORT"`
	isLoaded bool
}

var config Config

func loadConfig(configFilePath string, configName string, configType string) (Config, error) {
	viper.AddConfigPath(configFilePath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	// To read from program env variables also and override previous ones from config if exist
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}
	err = viper.Unmarshal(&config)
	if err == nil {
		config.isLoaded = true
	}
	return config, err
}

func GetConfig(configFilePath string) (Config, error) {
	if !config.isLoaded {
		return loadConfig(configFilePath, "prod", "env")
	}
	return config, nil
}