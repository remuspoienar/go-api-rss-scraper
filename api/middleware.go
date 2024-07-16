package api

import (
	"blogator/internal/database"
	"net/http"
	"regexp"
	"strings"
)

type authedHandler func(http.ResponseWriter, *http.Request, *database.User)

func (c *Config) RequireAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("Authorization")

		if !regexp.MustCompile(`^ApiKey \w{64}$`).MatchString(apiKey) {
			sendError(w, 401, "Invalid auth")
			return
		}

		apiKey = strings.Trim(apiKey, "ApiKey ")

		user, err := c.DB.FindUserByApiKey(r.Context(), apiKey)

		if err != nil {
			sendError(w, 404, "Cannot find user")
			return
		}
		handler(w, r, user)
	}

}
