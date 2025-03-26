package middleware

import (
	"context"
	"github.com/ITegs/trippify/utils"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
)

type contextKey string

const UserContextKey contextKey = "user"

type Auth struct {
}

func (Auth) Handler(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		tokenString := extractToken(r)
		if tokenString == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserContextKey, claims)
		next(w, r.WithContext(ctx), ps)
	}
}

func extractToken(r *http.Request) string {
	bearer := r.Header.Get("Authorization")
	if bearer == "" || !strings.HasPrefix(bearer, "Bearer ") {
		return ""
	}
	return strings.TrimPrefix(bearer, "Bearer ")
}
