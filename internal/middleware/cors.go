package middleware

import (
	"context"
	"log"
	"net/http"

	"example.com/bbb/internal/auth"
)

func CorsMiddleware(nxt http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		log.Println("Incoming request: Method -> ", r.Method, " URL -> ", r.URL.Path)
		log.Println("Origin: ", origin)

		if origin == "http://localhost:8080" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			log.Println("CORS allowed for origin", origin)
		} else {
			log.Println("CORS blocked for origin", origin)
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")

		if r.Method == http.MethodOptions {
			log.Println("Preflight OPTIONS request detected")
			w.WriteHeader(http.StatusOK)
			return
		}

		log.Println("Passing request to next handler")
		nxt(w, r)
	}
}

func AuthMiddleware(nxt http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		token, err := auth.VerifyIDToken(authHeader)
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		userID := token.UID
		if userID == "" {
			http.Error(w, "UserID not found in token", http.StatusBadRequest)
			return
		}

		log.Println("Successful Auth: User ->", userID)
		context := context.WithValue(r.Context(), "userID", userID)
		nxt(w, r.WithContext(context))
	}
}
