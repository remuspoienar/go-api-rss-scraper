package api

import (
	"blogator/internal/database"
)

type Config struct {
	DB                      *database.Queries
	FeedFetchConcurrency    int
	FeedFetchIntervalSecond int
}
