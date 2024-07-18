package main

import (
	"blogator/api"
	"blogator/internal/database"
	"blogator/scraping"
	"database/sql"
	_ "github.com/lib/pq"
)
import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("DB_URL"))
	dbQueries := database.New(db)
	config := &api.Config{
		DB:                      dbQueries,
		FeedFetchConcurrency:    3,
		FeedFetchIntervalSecond: 5,
	}

	go scraping.ScheduledFetchPosts(config)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/healthz", config.HealthCheck)
	mux.HandleFunc("GET /v1/err", config.Err)

	mux.HandleFunc("GET /v1/user", config.RequireAuth(config.GetUser))
	mux.HandleFunc("POST /v1/users", config.CreateUser)

	mux.HandleFunc("GET /v1/posts", config.RequireAuth(config.GetPosts))

	mux.HandleFunc("GET /v1/feeds", config.GetFeeds)
	mux.HandleFunc("POST /v1/feeds", config.RequireAuth(config.CreateFeed))
	mux.HandleFunc("POST /v1/feeds/{feedId}/follow", config.RequireAuth(config.FollowFeed))
	mux.HandleFunc("DELETE /v1/feeds/{feedId}/follow", config.RequireAuth(config.UnfollowFeed))

	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	fmt.Println("Server running on", addr)

	log.Fatal(http.ListenAndServe(addr, mux))

}
