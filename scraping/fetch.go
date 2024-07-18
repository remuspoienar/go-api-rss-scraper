package scraping

import (
	"blogator/internal/database"
	"encoding/xml"
	"fmt"
	"net/http"
)

type rootNode struct {
	XMLName xml.Name `xml:"rss"`
	Posts   []Post   `xml:"channel>item"`
}

type Post struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	PublishedAt string `xml:"pubDate"`
	Guid        string `xml:"guid"`
	Description string `xml:"description"`
}

func FetchFeed(feed *database.Feed) *[]Post {
	fmt.Println("Fetching feed", feed.Url)
	var posts []Post
	link := feed.Url

	resp, err := http.Get(link)

	if err != nil {
		fmt.Printf("Fetching %s ended in error %s\n", link, err.Error())

	}

	var parentNode rootNode
	err = xml.NewDecoder(resp.Body).Decode(&parentNode)

	if err != nil {
		fmt.Printf("Parsing content from %s ended in error %s\n", link, err.Error())
	}

	posts = parentNode.Posts

	fmt.Printf("Parsed feed %s , %v posts found\n", feed.Url, len(posts))
	return &posts
}
