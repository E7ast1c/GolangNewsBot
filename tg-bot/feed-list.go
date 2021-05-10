package tg_bot

import (
	"NewsFeedBot/tg-bot/news-provider"
	blog_golang_org "NewsFeedBot/tg-bot/news-provider/sources/blog-golang-org"
	golang_bridge "NewsFeedBot/tg-bot/news-provider/sources/golang-bridge"
	habr "NewsFeedBot/tg-bot/news-provider/sources/habr"
)

type FeedList struct {
	Url          string
	FeedProvider Provider
}

type Provider interface {
	GetNewsFeed(string) ([]news_provider.ResultFeed, error)
}

// GetFeedList storing list of channel subscriptions
func GetFeedList() map[string]FeedList {
	return map[string]FeedList{
		"habr": {Url: "https://habr.com/ru/rss/flows/develop/top/weekly/?fl=ru",
			FeedProvider: &habr.HabrFeed{}},
		"golangbridge": {Url: "https://forum.golangbridge.org/posts.rss",
			FeedProvider: &golang_bridge.GolangBridgeFeed{}},
		"bloggolangorg": {Url: "https://blog.golang.org/feed.atom?format=xml",
			FeedProvider: &blog_golang_org.BlogGolangOrgFeed{}},
	}
}