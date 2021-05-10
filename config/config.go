package config

import "github.com/spf13/viper"

type AppConfig struct {
	BotToken      string
	ServerAddress string
	WebhookUrl    string
}

func GetAppConfig() *AppConfig {
	viper.AutomaticEnv()
	token := viper.GetString("BOTTOKEN")
	webhook := viper.GetString("WEBHOOK")

	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.ReadInConfig()
	serverAddress := viper.GetString("server.port")

	return &AppConfig{
		BotToken:      token,
		ServerAddress: serverAddress,
		WebhookUrl:    webhook,
	}
}
