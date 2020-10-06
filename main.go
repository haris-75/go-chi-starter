package main

import (
	"fmt"
	"github.com/go-chi/jwtauth"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"

	"./apis"
	"./constants"
)

func main() {
	handleRequests()
}

func handleRequests() {
	r := chi.NewRouter()

	// Config
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	// CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Content-Type", "Jwt-Token"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Routes
	r.Group(protectedRoutes)
	r.Group(publicRoutes)

	//Run
	fmt.Printf("[START]\tStarting server on %v\n", constants.GetAPIAddress())
	log.Printf("[ERROR]\t%v\n", http.ListenAndServe(constants.GetHTTPPort(), r))
}

func publicRoutes(r chi.Router) {
	r.Get(constants.HomeAPIRoute, apis.HomeAPI)
	r.Post(constants.SignInAPIRoute, apis.SignIn)
}

func protectedRoutes(r chi.Router) {
	r.Use(jwtauth.Verifier(apis.TokenAuth))
	r.Use(jwtauth.Authenticator)

	r.Get(constants.AdminAPIRoute, apis.AdminAPI)
	r.Get(constants.UserAPIRoute, apis.UserAPI)
}
