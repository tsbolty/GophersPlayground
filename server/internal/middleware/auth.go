package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/tsbolty/GophersPlayground/internal/auth"
	"github.com/tsbolty/GophersPlayground/internal/redis"
)

type contextKey string

const userContextKey contextKey = "user"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Skip authentication for login and register operations
		if isAuthExempted(r) {
			next.ServeHTTP(w, r)
			return
		}

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		// Validate access token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := validateToken(tokenString)
		if err != nil {
			handleTokenError(w, err)
			return
		}

		userId := int(claims["id"].(float64))

		// Refresh the token if needed
		if shouldRefreshToken(claims) {
			if err := refreshTokenAndSetSession(w, userId); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		// Attach claims to context and continue
		ctx := context.WithValue(r.Context(), userContextKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func isAuthExempted(r *http.Request) bool {
	operationName, _ := getOperationNameFromRequest(r)
	return operationName == "login" || operationName == "register"
}

func validateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid or expired token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token claims")
}

func shouldRefreshToken(claims jwt.MapClaims) bool {
	exp := int64(claims["exp"].(float64))
	return exp < (time.Now().Unix() + 30)
}

func refreshTokenAndSetSession(w http.ResponseWriter, userId int) error {
	newAccessToken, newRefreshToken, err := auth.GenerateNewTokens(uint(userId))
	if err != nil {
		return err
	}

	// Update Redis session and response cookies/headers with new tokens
	updateSessionAndResponse(w, userId, newAccessToken, newRefreshToken)
	return nil
}

func updateSessionAndResponse(w http.ResponseWriter, userId int, accessToken, refreshToken string) {
	// Update Redis session with new access token
	redis.SetUserSession(userId, accessToken, refreshToken, time.Hour*24)

	// Set new tokens in response
	setRefreshTokenCookie(w, refreshToken)
	w.Header().Set("Authorization", "Bearer "+accessToken)
}

func setRefreshTokenCookie(w http.ResponseWriter, refreshToken string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(time.Hour * 24 * 7),
		HttpOnly: true,
		Path:     "/",
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})
}

type GraphQLRequestBody struct {
	OperationName string `json:"operationName"`
}

// getOperationNameFromRequest extracts the operation name from a GraphQL HTTP POST request
func getOperationNameFromRequest(r *http.Request) (string, error) {
	// Skip non-POST requests
	if r.Method != http.MethodPost {
		return "", nil
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	// Restore the body for future readers
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	// Check if body contains specific operations
	query := string(body)
	if strings.Contains(query, "mutation") && strings.Contains(query, "register") {
		return "register", nil
	}
	if strings.Contains(query, "mutation") && strings.Contains(query, "login") {
		return "login", nil
	}

	// Parse for operationName if present
	var reqBody GraphQLRequestBody
	if err := json.Unmarshal(body, &reqBody); err == nil && reqBody.OperationName != "" {
		return reqBody.OperationName, nil
	}

	// Default to not finding a specific operation
	return "", nil
}

func handleTokenError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusUnauthorized)
}
