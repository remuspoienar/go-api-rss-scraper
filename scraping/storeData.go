package scraping

import (
	"blogator/api"
	"blogator/internal"
	"blogator/internal/database"
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func StoreData(c *api.Config, feed *database.Feed, posts *[]Post) {
	err := c.DB.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		fmt.Println("Could not mark feed as fetched")
		return
	}

	var (
		savedCount    = 0
		existingCount = 0
	)
	for _, post := range *posts {
		existingPost, err := c.DB.GetPostByUrl(context.Background(), post.Link)
		if existingPost != nil {
			existingCount++
			continue
		}
		err = c.DB.CreatePost(context.Background(), &database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			PublishedAt: internal.ParseRFC1123ToTime(post.PublishedAt),
			Title:       post.Title,
			Url:         post.Link,
			Description: post.Description,
			FeedID:      feed.ID,
		})

		if err != nil {
			fmt.Println("Could not save post", post.Link, err.Error())
			continue
		}
		savedCount++
	}

	fmt.Printf("Found %d and saved %d posts in the db for feed %s\n", existingCount, savedCount, feed.Url)

}
