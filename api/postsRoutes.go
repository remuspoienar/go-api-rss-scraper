package api

import (
	"blogator/internal/database"
	"net/http"
	"strconv"
)

func (c *Config) GetPosts(w http.ResponseWriter, r *http.Request, user *database.User) {
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 50
	}

	posts, err := c.DB.GetPostsForUser(r.Context(), &database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})

	if len(posts) == 0 {
		posts = []*database.Post{}
	}

	sendJson(w, 200, posts)
}
