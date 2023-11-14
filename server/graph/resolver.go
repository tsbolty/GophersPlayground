package graph

// Other imports...
import (
	"context"
	"fmt"
	"strconv"

	"github.com/tsbolty/GophersPlayground/db/models/services"
	"github.com/tsbolty/GophersPlayground/db/models/todos"
	"github.com/tsbolty/GophersPlayground/graph/model"
)

type Resolver struct {
	ComplexService *services.ComplexService
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*todos.Todo, error) {
	userID, err := strconv.ParseUint(input.UserID, 10, 64)

	// On 32-bit systems, check for overflow
	if uint64(uint(userID)) != userID {
		// Handle the overflow, for example, return a GraphQL error
		return nil, fmt.Errorf("userID is too large for this platform")
	}

	// Now you can safely convert userID to uint
	todo, err := r.ComplexService.CreateTodoForUser(input.Text, uint(userID))

	if err != nil {
		return nil, err
	}

	return todo, nil
}

// func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
// 	todos, err := r.ComplexService.GetAllTodos()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return r.todos, nil
// }

// func (r *mutationResolver) CreateUser(ctx context.Context, name string, email string) (*model.User, error) {
// 	user, err := CreateUser(input.Name, input.Email)

// 	return user, nil
// }
