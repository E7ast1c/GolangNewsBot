package blog_golang_org

import (
	NewsProvider "NewsFeedBot/tg-bot/news-provider"
	"NewsFeedBot/transport"
	"encoding/xml"
	"errors"
)

type BlogGolangOrgFeed struct {
	XMLName xml.Name `xml:"feed"`
	Entry   []struct {
		Title string `xml:"title"`
		Link  struct {
			Href string `xml:"href,attr"`
		} `xml:"link"`
	} `xml:"entry"`
}


func (bgo *BlogGolangOrgFeed) GetNewsFeed(url string) ([]NewsProvider.ResultFeed, error) {
	body, err := transport.GetRSSFeed(url)
	if err != nil {
		return nil, err
	}

	if err = xml.Unmarshal(body, &bgo); err != nil {
		return nil, err
	}

	if bgo.Entry == nil && len(bgo.Entry) == 0 {
		return nil, errors.New("nil or empty result struct set")
	}

	rFeed := make([]NewsProvider.ResultFeed, 0, len(bgo.Entry))

	for i := 0; i < len(bgo.Entry); i++ {
		rFeed = append(rFeed, NewsProvider.ResultFeed{
			Title: bgo.Entry[i].Title,
			Url:   bgo.Entry[i].Link.Href,
		})
	}

	return rFeed, nil
}
