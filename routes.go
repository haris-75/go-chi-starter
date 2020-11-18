package main

import (
	"./apis"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

func publicRoutes(r chi.Router) {
	r.Get("/", apis.HomeAPI)
	r.Post("/login", apis.SignIn)
}

func protectedRoutes(r chi.Router) {
	r.Use(jwtauth.Verifier(apis.TokenAuth))
	r.Use(authenticator)

	r.Get("/admin", apis.AdminAPI)
	r.Get("/user", apis.UserAPI)
}
