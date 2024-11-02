package middlewares

import (
	"backend/internal/utils"
	"context"
	"net/http"
	"strings"
)

func Auth(jwtUtil *utils.JWTUtil) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			token := r.Header.Get("Authorization")
			if token == "" {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			if !strings.HasPrefix(token, "Bearer ") {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			token = strings.TrimPrefix(token, "Bearer ")

			claims, err := jwtUtil.ValidateToken(token)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), utils.UserKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
