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
		sendError(w, 422, "Cannot create feed "+err.Error())
		return
	}

	sendJson(w, 201, feed)

}

func (c *Config) GetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := c.DB.GetFeeds(r.Context())
	if err != nil {
		sendError(w, 422, "Cannot get feeds "+err.Error())
		return
	}

	sendJson(w, 200, feeds)

}
