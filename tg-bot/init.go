package tg_bot

import (
	"NewsFeedBot/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
)

const (
	WebhookURL = "https://9e0c93e98738.ngrok.io"
)

func StartBot(config config.AppConfig) (*tgbotapi.BotAPI, error) {
	bot, err := tgbotapi.NewBotAPI(config.BotToken)
	if err != nil {
		return nil, err
	}

	log.Printf("Bot authorized on account %s", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhook(WebhookURL))
	if err != nil {
		return nil, err
	}

	return bot, err
}
