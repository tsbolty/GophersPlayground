package services

import (
	"github.com/tsbolty/GophersPlayground/db/models/todos"
	"github.com/tsbolty/GophersPlayground/db/models/users"
)

type ComplexService struct {
	TodoRepo todos.TodoRepository
	UserRepo users.UserRepository
}

func NewComplexService(userRepo users.UserRepository, todoRepo todos.TodoRepository) *ComplexService {
	return &ComplexService{
		TodoRepo: todoRepo,
		UserRepo: userRepo,
	}
}

func (s *ComplexService) CreateTodoForUser(text string, userID uint) (*todos.Todo, error) {
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

func (s *ComplexService) GetTodosForUser(userID uint) ([]*todos.Todo, error) {
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
