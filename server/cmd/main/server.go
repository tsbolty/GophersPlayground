package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/playground"
	db "github.com/tsbolty/GophersPlayground/cmd/main/db"
	auth "github.com/tsbolty/GophersPlayground/internal/auth"
	"github.com/tsbolty/GophersPlayground/internal/middleware"
	"github.com/tsbolty/GophersPlayground/internal/redis"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dbInstance, dbErr := db.InitializeDB()
	if dbErr != nil {
		log.Fatalf("Could not connect to db: %v", dbErr)
	}

	sqlDB, err := dbInstance.DB()
	if err != nil {
		log.Fatalf("Could not get sql.DB: %v", err)
	}
	defer sqlDB.Close()

	// Initialize Redis
	redis.InitializeRedis()

	srv := InitializeServices(dbInstance)

	// Wrap the srv with the AuthMiddleware
	authedGraphQLServer := middleware.AuthMiddleware(srv)

	// handle token refresh
	http.HandleFunc("/api/token/refresh", auth.RefreshTokenHandler)

	http.Handle("/", playground.Handler("GraphQL playground", "/api/graphql"))
	http.Handle("/api/graphql", authedGraphQLServer)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
