package tg_bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	log "github.com/sirupsen/logrus"
)

func ListenAndHandleUpdates(bot *tgbotapi.BotAPI) {
	defer func() {
		if r := recover(); r != nil {
			log.Error("panic un recovered")
		}
	}()

	feedList := GetFeedList()
	updates := bot.ListenForWebhook("/")

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Infof("received message = %s, from user = %s\n", update.Message.Text, update.Message.From.UserName)

		comm := update.Message.Command()
		if comm == "" {
			if _, err := bot.Send(
				tgbotapi.NewMessage(update.Message.Chat.ID, "Use commands, not words\n")); err != nil {
				log.Error(err)
			}
			continue
		}

		// send greeting message
		if comm == "start" {
			if err := greeting(bot, update); err != nil {
				log.Errorf("error on send greeting msg, err = %s\n",err)
			}
			continue
		}

		if feed, ok := feedList[comm]; ok {
			list, err := feed.FeedProvider.GetNewsFeed(feed.Url)
			if err != nil {
				log.Errorf("error on receive news by command = %s, url = %s, err = %s\n", comm, feed.Url, err)
				continue
			}

			for _, article := range list {
				if _, err = bot.Send(
					tgbotapi.NewMessage(update.Message.Chat.ID, article.Title+"\n"+article.Url)); err != nil {
					log.Error(err)
					continue
				}
			}
		} else {
			if _, err := bot.Send(
				tgbotapi.NewMessage(update.Message.Chat.ID, "I can`t find this command")); err != nil {
				log.Error(err)
				continue
			}
		}
	}
}

// greeting send custom greeting message
func greeting (bot *tgbotapi.BotAPI, update tgbotapi.Update) error {
	var feedText string
	for k,_ := range GetFeedList() {
		feedText += "/"+k+"\n"
	}

	text := fmt.Sprintf("Hello, you can receive the following news feeds\n%s", feedText)
	if _, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, text)); err != nil {
		return err
	}
	return nil
}