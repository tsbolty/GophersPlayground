package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	users "github.com/tsbolty/GophersPlayground/internal/user"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserService          *users.UserService
	GenerateNewTokens    func(userID uint) (string, string, error)
	GenerateAccessToken  func(userId uint) (string, error)
	GenerateRefreshToken func(userId uint) (string, error)
	ValidateRefreshToken func(token string) (uint, error)
	RefreshTokenHandler  func(w http.ResponseWriter, r *http.Request)
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

func RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	refreshTokenCookie, err := r.Cookie("auth_token")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(w, "Refresh token not provided", http.StatusUnauthorized)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	validatedRefreshToken, err := ValidateRefreshToken(refreshTokenCookie.Value)
	if err != nil {
		http.Error(w, "Failed to parse token", http.StatusInternalServerError)
		return
	}

	accessToken, newRefreshToken, err := GenerateNewTokens(validatedRefreshToken)
	if err != nil {
		http.Error(w, "Failed to generate new tokens", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    newRefreshToken,
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
	})

	// Send the new access token and refresh token back in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"accessToken": accessToken, "refreshToken": newRefreshToken})
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
