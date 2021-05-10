package config

import "github.com/spf13/viper"

type AppConfig struct {
	BotToken string
	ServerAddress string
}

func GetAppConfig() *AppConfig {
	viper.AutomaticEnv()
	token := viper.GetString("BOTTOKEN")

	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.ReadInConfig()
	serverAddress := viper.GetString("server.port")

	return &AppConfig{
		BotToken:     token,
		ServerAddress: serverAddress,
	}
}
