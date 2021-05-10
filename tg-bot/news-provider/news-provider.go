package news_provider

type ResultFeed struct {
	Title string `mapper:"title"`
	Url   string `mapper:"url"`
}
