package auth

import (
	"fmt"

	users "github.com/tsbolty/GophersPlayground/internal/user"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserService *users.UserService
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"user"`
}

func NewAuthService(userService *users.UserService) *AuthService {
	return &AuthService{
		UserService: userService,
	}
}

func (a *AuthService) AuthenticateUser(email string, password string) (token string, user *users.User, err error) {
	user, err = a.UserService.FindUserByEmail(email)
	if err != nil {
		return "", nil, err // User not found or other error
	}

	// Verify password
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", nil, err // Invalid password
	}

	token, err = GenerateToken(user.ID)
	if err != nil {
		return "", nil, err // Error generating token
	}

	return token, user, nil
}

func (a *AuthService) RegisterUser(email string, name string, password string) (token string, user *users.User, err error) {
	fmt.Println("RegisterUser")
	// Hash password
	hashedPassword, err := hashPassword(password)
	fmt.Println("hashedPassword", hashedPassword)
	if err != nil {
		return "", nil, err
	}

	user, err = a.UserService.CreateUser(email, name, hashedPassword)
	fmt.Println("CREATED USER", user)
	if err != nil {
		return "", nil, err
	}

	token, err = GenerateToken(user.ID)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
