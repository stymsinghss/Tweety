package handler

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

// KeyAuthUserId to use in context
const KeyAuthUserId key = "auth_user_id"
type key string

// withAuth decodes the token and extract the user data from it
func (h *handler) withAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// token comes from Authorization header and has the prefix of "Bearer"
		a := r.Header.Get("Authorization")

		// Do nothing if "Bearer Token" is not present
		if !strings.HasPrefix(a,"Bearer ") {
			next.ServeHTTP(w,r)
			return
		}

		// extract token after "Bearer"
		token := a[7:]
		uid, err := h.AuthUserId(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// set user_id to context so that other routes can use it
		ctx := r.Context()
		ctx = context.WithValue(ctx, KeyAuthUserId, uid)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}