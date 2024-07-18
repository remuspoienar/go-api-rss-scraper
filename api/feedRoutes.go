package api

import (
	"blogator/internal/database"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type feedPayload struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (c *Config) GetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := c.DB.GetFeeds(r.Context())
	if err != nil {
		sendError(w, 422, "Cannot get feeds "+err.Error())
		return
	}

	sendJson(w, 200, feeds)
}

func (c *Config) CreateFeed(w http.ResponseWriter, r *http.Request, user *database.User) {
	var feedInput feedPayload
	err := json.NewDecoder(r.Body).Decode(&feedInput)

	if err != nil {
		sendError(w, 400, "Invalid feed payload")
		return
	}

	feed, err := c.DB.CreateFeed(r.Context(), &database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      feedInput.Name,
		Url:       feedInput.Url,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID})

	if err != nil {
		sendError(w, 400, "Cannot create feed "+err.Error())
		return
	}

	follow, _ := c.DB.FollowFeed(r.Context(), &database.FollowFeedParams{
		FeedID:     feed.ID,
		UserID:     user.ID,
		FollowedAt: time.Now().UTC()})

	sendJson(w, 201, map[string]any{"feed": feed, "follow": follow})

}

func (c *Config) FollowFeed(w http.ResponseWriter, r *http.Request, user *database.User) {
	feedId := r.PathValue("feedId")

	if len(feedId) == 0 {
		sendError(w, 400, "Invalid feed payload")
		return
	}

	feedFollow, err := c.DB.FollowFeed(r.Context(), &database.FollowFeedParams{
		FeedID:     uuid.MustParse(feedId),
		UserID:     user.ID,
		FollowedAt: time.Now().UTC()})

	if err != nil {
		sendError(w, 400, "Cannot create feed "+err.Error())
		return
	}

	sendJson(w, 201, feedFollow)
}

func (c *Config) UnfollowFeed(w http.ResponseWriter, r *http.Request, user *database.User) {
	feedId := r.PathValue("feedId")

	if len(feedId) == 0 {
		sendError(w, 400, "Invalid feed payload")
		return
	}

	feedFollow, err := c.DB.GetFeedFollow(r.Context(), &database.GetFeedFollowParams{
		FeedID: uuid.MustParse(feedId),
		UserID: user.ID,
	})

	if err != nil {
		sendError(w, 400, "This feed is not followed")
		return
	}

	c.DB.UnfollowFeed(r.Context(), &database.UnfollowFeedParams{
		FeedID: feedFollow.FeedID,
		UserID: feedFollow.UserID,
	})

	sendJson(w, 204, nil)
}
