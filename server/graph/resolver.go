package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"github.com/tsbolty/GophersPlayground/internal/auth"
	"github.com/tsbolty/GophersPlayground/internal/services"
	todos "github.com/tsbolty/GophersPlayground/internal/todo"
	users "github.com/tsbolty/GophersPlayground/internal/user"
)

type Resolver struct {
	AuthService         *auth.AuthService
	ComplexService      *services.ComplexService
	TodoService         *services.TodoService
	UserService         *services.UserService
	UserBusinessService *users.UserService
	TodoBusinessService *todos.TodoService
}
