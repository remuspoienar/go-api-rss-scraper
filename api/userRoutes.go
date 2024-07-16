package api

import (
	"blogator/internal/database"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type userPayload struct {
	Name string `json:"name"`
}

func (c *Config) CreateUser(w http.ResponseWriter, r *http.Request) {
	var u userPayload
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		sendError(w, 400, "Invalid user payload")
		return
	}

	user, err := c.DB.CreateUser(r.Context(), &database.CreateUserParams{
		ID:        uuid.New(),
		Name:      u.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC()})

	if err != nil {
		sendError(w, 422, "Cannot create user")
		return
	}

	sendJson(w, 201, user)

}

func (c *Config) GetUser(w http.ResponseWriter, r *http.Request, user *database.User) {
	sendJson(w, 201, *user)

}
