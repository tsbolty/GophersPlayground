package users

import (
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *User) error
	FindByID(id uint) (*User, error)
	FindByEmail(email string, user *User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(user *User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByID(id uint) (*User, error) {
	var user User
	result := r.db.First(&user, id)
	return &user, result.Error
}

func (r *userRepository) FindByEmail(email string, user *User) error {
	result := r.db.Where("email = ?", email).First(user)
	return result.Error // This will return gorm.ErrRecordNotFound if the user is not found
}
