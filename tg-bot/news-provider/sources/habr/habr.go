package habr

import (
	NewsProvider "NewsFeedBot/tg-bot/news-provider"
	"NewsFeedBot/transport"
	"encoding/xml"
	"errors"
)

type HabrFeed struct {
	XMLName xml.Name `xml:"rss"`
	Channel struct {
		Item []struct {
			Title string `xml:"title"`
			Link  string `xml:"link"`
		} `xml:"item"`
	} `xml:"channel"`
}

func (h *HabrFeed) GetNewsFeed(url string) ([]NewsProvider.ResultFeed, error) {
	body, err := transport.GetRSSFeed(url)
	if err != nil {
		return nil, err
	}

	if err = xml.Unmarshal(body, &h); err != nil {
		return nil, err
	}

	if h.Channel.Item == nil && len(h.Channel.Item) == 0 {
		return nil, errors.New("nil or empty result struct set")
	}

	rFeed := make([]NewsProvider.ResultFeed, 0, len(h.Channel.Item))

	for i := 0; i < len(h.Channel.Item); i++ {
		rFeed = append(rFeed, NewsProvider.ResultFeed{
			Title: h.Channel.Item[i].Title,
			Url:   h.Channel.Item[i].Link,
		})
	}

	return rFeed, nil
}
