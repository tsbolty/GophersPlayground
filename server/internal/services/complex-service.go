package services

import (
	todos "github.com/tsbolty/GophersPlayground/internal/todo"
	users "github.com/tsbolty/GophersPlayground/internal/user"
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
