package users

import (
	"errors"
)

type UserService interface {
	CreateUser(email, name string) (*User, error)
	GetUserByID(id uint) (*User, error)
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) CreateUser(email, name string) (*User, error) {
	var existingUser User
	if err := s.repo.FindByEmail(email, &existingUser); err == nil {
		return nil, errors.New("user already exists")
	}

	user := &User{
		Name:  name,
		Email: email,
	}

	err := s.repo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetUserByID(id uint) (*User, error) {
	return s.repo.FindByID(id)
}
