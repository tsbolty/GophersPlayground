package users

import (
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(email string, name string, password string) (*User, error)
	FindByID(id uint) (*User, error)
	FindByEmail(email string) (*User, error)
	FindAll() ([]*User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(email string, name string, password string) (*User, error) {
	user := &User{
		Email:    email,
		Name:     name,
		Password: password,
	}
	result := r.db.Create(&user)
	return user, result.Error
}

func (r *userRepository) FindByID(id uint) (*User, error) {
	var user User
	result := r.db.First(&user, id)
	return &user, result.Error
}

func (r *userRepository) FindByEmail(email string) (*User, error) {
	var user User
	result := r.db.First(&user, "email = ?", email)

	return &user, result.Error
}

func (r *userRepository) FindAll() ([]*User, error) {
	var users []*User
	result := r.db.Find(&users)
	return users, result.Error
}
