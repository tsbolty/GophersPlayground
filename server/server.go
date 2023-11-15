package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/tsbolty/GophersPlayground/db/models/services"
	"github.com/tsbolty/GophersPlayground/db/models/todos"
	"github.com/tsbolty/GophersPlayground/db/models/users"
	"github.com/tsbolty/GophersPlayground/graph"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dbInstance, dbErr := InitializeDB()
	if dbErr != nil {
		log.Fatalf("Could not connect to db: %v", dbErr)
	}

	sqlDB, err := dbInstance.DB()
	if err != nil {
		log.Fatalf("Could not get sql.DB: %v", err)
	}
	defer sqlDB.Close()

	// store := NewStore(dbInstance)

	// Initialize repositories
	userRepo := users.NewUserRepository(dbInstance)
	todoRepo := todos.NewTodoRepository(dbInstance)

	// Initialize services
	userService := users.NewUserService(userRepo)
	todoService := todos.NewTodoService(todoRepo)
	complexService := services.NewComplexService(userRepo, todoRepo)

	// Set up GraphQL server with all services
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			UserService:    userService,
			TodoService:    todoService,
			ComplexService: complexService,
		},
	}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
