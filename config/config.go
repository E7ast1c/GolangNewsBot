package config

import "github.com/spf13/viper"

type AppConfig struct {
	BotToken   string
	ServerPort string
	WebhookUrl string
}

func GetAppConfig() *AppConfig {
	viper.AutomaticEnv()
	token := viper.GetString("BOTTOKEN")
	webhook := viper.GetString("WEBHOOK")
	port := viper.GetString("PORT")

	return &AppConfig{
		BotToken:   token,
		ServerPort: port,
		WebhookUrl: webhook,
	}
}
