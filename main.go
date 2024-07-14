package main

import (
	"blogator/api"
	"blogator/internal/database"
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
	config := api.Config{
		DB: dbQueries,
	}
	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/healthz", config.HealthCheck)
	mux.HandleFunc("GET /v1/err", config.Err)

	mux.HandleFunc("POST /v1/users", config.CreateUser)

	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	fmt.Println("Server running on", addr)
	log.Fatal(http.ListenAndServe(addr, mux))

}
