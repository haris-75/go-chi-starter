package main

import (
	"./log"
	"github.com/go-chi/jwtauth"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"

	"./constants"
	"./utils"
)

func main() {
	handleRequests()
}

func handleRequests() {
	r := chi.NewRouter()

	// Config
	r.Use(middleware.RequestID)
	r.Use(log.RequestLogger)
	r.Use(log.RequestFileLogger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	// CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Access-Control-Allow-Origin"},
		ExposedHeaders:   []string{"Content-Type", "JWT-Token"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Routes
	r.Group(protectedRoutes)
	r.Group(publicRoutes)

	//Run
	log.Info.Printf("Starting server on %v\n", constants.GetAPIAddress())
	log.Error.Printf("%v\n", http.ListenAndServe(constants.GetHTTPPort(), r))
}

// Authenticator  middleware
func authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _, err := jwtauth.FromContext(r.Context())

		if err != nil || token == nil || !token.Valid {
			utils.RespondError(w, http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
