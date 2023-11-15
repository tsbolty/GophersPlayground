package graph

import (
	"github.com/tsbolty/GophersPlayground/db/models/services"
	"github.com/tsbolty/GophersPlayground/db/models/todos"
	"github.com/tsbolty/GophersPlayground/db/models/users"
)

type Resolver struct {
	ComplexService *services.ComplexService
	UserService    *users.UserService
	TodoService    *todos.TodoService
}
