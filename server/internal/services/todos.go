package services

import (
	todos "github.com/tsbolty/GophersPlayground/internal/todo"
	users "github.com/tsbolty/GophersPlayground/internal/user"
)

type TodoService struct {
	TodoRepo todos.TodoRepository
	UserRepo users.UserRepository
}

func NewTodoService(userRepo users.UserRepository, todoRepo todos.TodoRepository) *TodoService {
	return &TodoService{
		TodoRepo: todoRepo,
		UserRepo: userRepo,
	}
}

func (s *TodoService) CreateTodoForUser(text string, userID uint) (*todos.Todo, error) {
	user, err := s.UserRepo.FindByID(userID) // Assuming FindByID is a method on UserRepository
	if err != nil {
		return nil, err
	}

	newTodo := &todos.Todo{
		Text:   text,
		Done:   false,
		UserID: user.ID,
	}

	createdTodo, err := s.TodoRepo.Create(newTodo)
	if err != nil {
		return nil, err
	}

	return createdTodo, nil
}

func (s *TodoService) GetTodosForUser(userID uint) ([]*todos.Todo, error) {
	foundUser, err := s.UserRepo.FindByID(userID)

	if err != nil {
		return nil, err
	}

	todos, err := s.TodoRepo.FindAllByUserID(foundUser.ID)

	if err != nil {
		return nil, err
	}

	return todos, nil
}
