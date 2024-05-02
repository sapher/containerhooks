package internal

import "github.com/spf13/viper"

type AppConfigKey string

type AppConfig struct{}

func LoadConfig() (*AppConfig, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var appConfig AppConfig
	err = viper.Unmarshal(&appConfig)
	return &appConfig, err
}
