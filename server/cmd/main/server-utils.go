package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/tsbolty/GophersPlayground/graph"
	"github.com/tsbolty/GophersPlayground/internal/services"
	todos "github.com/tsbolty/GophersPlayground/internal/todo"
	users "github.com/tsbolty/GophersPlayground/internal/user"
	"gorm.io/gorm"
)

// InitializeServices initializes all services
func InitializeServices(dbInstance *gorm.DB) *handler.Server {

	// Initialize repositories
	userRepo := users.NewUserRepository(dbInstance)
	todoRepo := todos.NewTodoRepository(dbInstance)

	// Initialize services
	userBusinessService := users.NewUserService(userRepo)
	todoBusinessService := todos.NewTodoService(todoRepo)
	complexService := services.NewComplexService(userRepo, todoRepo)
	todoService := services.NewTodoService(userRepo, todoRepo)
	userService := services.NewUserService()

	// Set up GraphQL server with all services
	return handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			UserBusinessService: userBusinessService,
			TodoBusinessService: todoBusinessService,
			ComplexService:      complexService,
			TodoService:         todoService,
			UserService:         userService,
		},
	}))
}
