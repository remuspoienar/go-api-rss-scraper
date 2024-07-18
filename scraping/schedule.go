package scraping

import (
	"blogator/api"
	"context"
	"fmt"
	"sync"
	"time"
)

func ScheduledFetchPosts(c *api.Config) {
	ticker := time.NewTicker(time.Second * time.Duration(c.FeedFetchIntervalSecond))

	for {
		select {
		case <-ticker.C:
			fmt.Print("FETCHING ------------>\n")
			feeds, _ := c.DB.GetNextFeedsToFetch(context.Background(), int32(c.FeedFetchConcurrency))
			wg := sync.WaitGroup{}
			for _, feed := range feeds {
				wg.Add(1)
				go func() {
					defer wg.Done()
					posts := FetchFeed(feed)
					StoreData(c, feed, posts)
				}()
			}

			wg.Wait()
			fmt.Print("FINISHED <------------\n\n")
		}
	}
}
