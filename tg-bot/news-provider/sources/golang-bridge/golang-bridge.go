package golang_bridge

import (
	NewsProvider "NewsFeedBot/tg-bot/news-provider"
	"NewsFeedBot/transport"
	"encoding/xml"
	"errors"
)

type GolangBridgeFeed struct {
	XMLName xml.Name `xml:"rss"`
	Channel struct {
		Item []struct {
			Title string `xml:"title"`
			Link  string `xml:"link"`
		} `xml:"item"`
	} `xml:"channel"`
}

func (gb *GolangBridgeFeed) GetNewsFeed(url string) ([]NewsProvider.ResultFeed, error) {
	body, err := transport.GetRSSFeed(url)
	if err != nil {
		return nil, err
	}

	if err = xml.Unmarshal(body, &gb); err != nil {
		return nil, err
	}

	if gb.Channel.Item == nil && len(gb.Channel.Item) == 0 {
		return nil, errors.New("nil or empty result struct set")
	}

	rFeed := make([]NewsProvider.ResultFeed, 0, len(gb.Channel.Item))

	for i := 0; i < len(gb.Channel.Item); i++ {
		rFeed = append(rFeed, NewsProvider.ResultFeed{
			Title: gb.Channel.Item[i].Title,
			Url:   gb.Channel.Item[i].Link,
		})
	}

	return rFeed, nil
}
