package main

import (
	"datumbrain/my-project/log"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
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
	httpPort := fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))

	log.Info.Printf("Starting server on %v\n", httpPort)
	log.Error.Println(http.ListenAndServe(httpPort, r))
}
