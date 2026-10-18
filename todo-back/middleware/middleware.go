package middleware

import (
	"context"
	"go-learning/todoApp/auth"
	"go-learning/todoApp/response"
	"net/http"
	"strings"
)

type contextKey string

const UserContextKey contextKey = "user"

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		token := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := auth.VerifyAccessToken(token)
		if err != nil {
			response.WriteErrMessage(w, http.StatusUnauthorized, "unauthorized")
			return
		}

		ctx := context.WithValue(r.Context(), UserContextKey, claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
