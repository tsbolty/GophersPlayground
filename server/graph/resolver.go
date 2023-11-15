package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	todos "github.com/tsbolty/GophersPlayground/internal/todo"
	users "github.com/tsbolty/GophersPlayground/internal/user"
	"github.com/tsbolty/GophersPlayground/pkg/services"
)

type Resolver struct {
	ComplexService      *services.ComplexService
	TodoService         *services.TodoService
	UserService         *services.UserService
	UserBusinessService *users.UserService
	TodoBusinessService *todos.TodoService
}
