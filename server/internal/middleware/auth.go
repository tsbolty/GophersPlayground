package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type contextKey string

const userContextKey contextKey = "user"

// AuthMiddleware checks for the presence of an Authorization header and validates the JWT token
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/api/graphql") {
			operationName, err := getOperationNameFromRequest(r)
			if err != nil {
				fmt.Println("Error reading request body:", err)
				http.Error(w, "Error reading request body", http.StatusInternalServerError)
				return
			}
			if operationName == "login" || operationName == "register" {
				fmt.Println("Skipping auth for operation:", operationName)
				next.ServeHTTP(w, r)
				return
			}
		}

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Return the key for validation
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Extract claims and attach to context
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx := context.WithValue(r.Context(), userContextKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
		}
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

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	// Restore the body for future readers
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

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
