package main

import (
	"datumbrain/my-project/apis"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

func publicRoutes(r chi.Router) {
	r.Get("/", apis.HomeAPI)
	r.Post("/login", apis.SignIn)
}

func protectedRoutes(r chi.Router) {
	r.Use(jwtauth.Verifier(apis.TokenAuth))
	r.Group(adminRoutes)
	r.Group(userRoutes)
}

func adminRoutes(r chi.Router) {
	r.Use(apis.Authenticator)
	r.Get("/admin", apis.AdminAPI)
}

func userRoutes(r chi.Router) {
	r.Use(apis.Authenticator)
	r.Get("/user", apis.UserAPI)
}
