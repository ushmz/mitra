package main

import (
	"context"
	"strings"

	"net/http"

	firebase "firebase.google.com/go"
	"github.com/rs/cors"
)

var (
	app firebase.App
)

func corsHandler() func(http.Handler) http.Handler {
	opt := cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}
	handler := cors.New(opt).Handler
	return handler
}

func adminOnly() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			isAdmin, ok := r.Context().Value("acl.admin").(bool)
			if !ok || !isAdmin {
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

// firebaseAuth : Auth middleware that check "Authorization" header and verify token
func firebaseAuth(app *firebase.App) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			if ctx == nil {
				ctx = context.Background()
			}

			client, err := app.Auth(ctx)
			if err != nil {
				// errors.New("Failed to get auth client")
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				return
			}

			auth := r.Header.Get("Authorization")
			idToken := strings.Replace(auth, "Bearer ", "", 1)

			if _, err = client.VerifyIDTokenAndCheckRevoked(ctx, idToken); err != nil {
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
