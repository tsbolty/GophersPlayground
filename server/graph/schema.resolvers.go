package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"fmt"
	"strconv"

	"github.com/tsbolty/GophersPlayground/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	userID, err := strconv.ParseUint(input.UserID, 10, 64)

	// On 32-bit systems, check for overflow
	if uint64(uint(userID)) != userID {
		// Handle the overflow, for example, return a GraphQL error
		return nil, fmt.Errorf("userID is too large for this platform")
	}

	todo, err := r.TodoService.CreateTodoForUser(input.Text, uint(userID))

	if err != nil {
		return nil, err
	}

	graphqlTodo := &model.Todo{
		ID:     fmt.Sprintf("%d", todo.ID),
		Text:   todo.Text,
		Done:   todo.Done,
		UserID: int(todo.UserID),
	}

	return graphqlTodo, nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.LoginUser) (*model.AuthPayload, error) {
	token, user, err := r.AuthService.AuthenticateUser(input.Email, input.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to authenticate user")
	}

	return &model.AuthPayload{
		Token: token,
		User: &model.User{
			ID:    fmt.Sprintf("%d", user.ID),
			Name:  user.Name,
			Email: user.Email,
		},
	}, nil
}

// Register is the resolver for the register field.
func (r *mutationResolver) RegisterUser(ctx context.Context, input model.NewUser) (*model.AuthPayload, error) {
	token, dbUser, err := r.AuthService.RegisterUser(input.Email, input.Name, input.Password)
	if err != nil {
		return nil, err
	}
	return &model.AuthPayload{
		Token: token,
		User: &model.User{
			ID:    fmt.Sprintf("%d", dbUser.ID),
			Name:  dbUser.Name,
			Email: dbUser.Email,
		},
	}, nil
}

// FindAllUsers is the resolver for the findAllUsers field.
func (r *queryResolver) FindAllUsers(ctx context.Context) ([]*model.User, error) {
	dbUsers, err := r.UserBusinessService.FindAllUsers()
	if err != nil {
		return nil, err
	}

	graphqlUsers := make([]*model.User, len(dbUsers))

	for i, user := range dbUsers {
		graphqlUsers[i] = &model.User{
			ID:    fmt.Sprintf("%d", user.ID),
			Name:  user.Name,
			Email: user.Email,
		}
	}

	return graphqlUsers, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	todos, err := r.TodoBusinessService.GetAllTodos()

	if err != nil {
		return nil, err
	}

	graphqlTodos := make([]*model.Todo, len(todos))

	for i, todo := range todos {
		graphqlTodos[i] = &model.Todo{
			ID:     fmt.Sprintf("%d", todo.ID),
			Text:   todo.Text,
			Done:   todo.Done,
			UserID: int(todo.UserID),
		}
	}

	return graphqlTodos, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
