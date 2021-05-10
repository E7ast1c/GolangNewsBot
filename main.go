package main

import (
	"NewsFeedBot/config"
	tgbot "NewsFeedBot/tg-bot"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	config := config.GetAppConfig()

	go func() {
		servErr := http.ListenAndServe(":"+config.ServerPort, nil)
		if servErr != nil {
			log.Fatal(servErr)
		}
	}()

	tgBot, err := tgbot.StartBot(*config)
	if err != nil {
		log.Fatalf("starting bot error = %s", err)
	}

	tgbot.ListenAndHandleUpdates(tgBot)
}