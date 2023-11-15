package graph

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
