package auth

import (
	"fmt"

	users "github.com/tsbolty/GophersPlayground/internal/user"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserService          *users.UserService
	GenerateNewTokens    func(userID uint) (string, string, error)
	GenerateAccessToken  func(userId uint) (string, error)
	GenerateRefreshToken func(userId uint) (string, error)
	ValidateRefreshToken func(token string) (uint, error)
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

func (a *AuthService) AuthenticateUser(email string, password string) (accessToken string, refreshToken string, user *users.User, err error) {
	user, err = a.UserService.FindUserByEmail(email)
	if err != nil {
		return "", "", nil, err // User not found or other error
	}

	// Verify password
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", "", nil, err // Invalid password
	}

	accessToken, refreshToken, tokenErr := GenerateNewTokens(user.ID)
	if tokenErr != nil {
		return "", "", nil, err // Error generating token
	}

	return accessToken, refreshToken, user, nil
}

func (a *AuthService) RegisterUser(email string, name string, password string) (toaccessToken string, refreshToken string, user *users.User, err error) {
	fmt.Println("RegisterUser")
	// Hash password
	hashedPassword, err := hashPassword(password)
	fmt.Println("hashedPassword", hashedPassword)
	if err != nil {
		return "", "", nil, err
	}

	user, err = a.UserService.CreateUser(email, name, hashedPassword)
	if err != nil {
		return "", "", nil, err
	}

	accessToken, refreshToken, tokenErr := GenerateNewTokens(user.ID)
	if tokenErr != nil {
		return "", "", nil, err // Error generating token
	}

	return accessToken, refreshToken, user, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
