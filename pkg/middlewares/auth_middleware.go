package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/namanag0502/go-todo-server/pkg/models"
	"github.com/namanag0502/go-todo-server/pkg/utils"
)

const (
	BearerTokenPrefix = "Bearer "
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if !strings.HasPrefix(header, BearerTokenPrefix) {
			http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
			return
		}
		token := strings.TrimPrefix(header, BearerTokenPrefix)
		claims, err := utils.VerifyToken(token)
		if err != nil {
			http.Error(w, "Failed to verify access token", http.StatusUnauthorized)
			return
		}

		userID := claims.Subject
		ctx := context.WithValue(r.Context(), models.UserContextKey, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
