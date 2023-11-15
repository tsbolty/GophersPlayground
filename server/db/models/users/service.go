package users

import (
	"errors"
	"log"
)

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(email string, name string) (*User, error) {
	foundUser, err := s.repo.FindByEmail(email)

	// Todo: Check if error is an actual db operation error

	if foundUser != nil {
		return nil, errors.New("user already exists")
	}

	user, err := s.repo.Create(email, name)
	if err != nil {
		return nil, err
	}

	log.Println("ABOUT TO RETURN USER FROM SERVICE")

	return user, nil
}

func (s *UserService) GetUserByID(id uint) (*User, error) {
	return s.repo.FindByID(id)
}
