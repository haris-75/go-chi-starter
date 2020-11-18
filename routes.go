package main

import (
	"./apis"
	"./log"
	"./models"
	"./utils"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"net/http"
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
	r.Use(adminAuthenticator)

	r.Get("/admin", apis.AdminAPI)
}

func userRoutes(r chi.Router) {
	r.Use(userAuthenticator)

	r.Get("/user", apis.UserAPI)
}

func adminAuthenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		user, err := getUserFromRequest(r)
		if err != nil {
			log.Warn.Println("Someone tried to access protected method without authentication.")
			utils.RespondCustomError(w, http.StatusUnauthorized, "Authorization information missing/invalid.")
			return
		}

		if user.Role != models.USER_ADMIN {
			log.Warn.Printf("`%v` tried to access protected method.\n", user.Name)
			utils.RespondCustomError(w, http.StatusUnauthorized, "Only admin can access.")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func userAuthenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		_, err := getUserFromRequest(r)
		if err != nil {
			log.Warn.Println("Someone tried to access protected method without authentication.")
			utils.RespondCustomError(w, http.StatusUnauthorized, "Authorization information missing/invalid.")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func getUserFromRequest(r *http.Request) (models.User, error) {
	token, claims, err := jwtauth.FromContext(r.Context())

	if err != nil || token == nil || !token.Valid {
		log.Error.Printf("%v\n", err)
		return models.User{}, fmt.Errorf("Authorization information missing/invalid.")
	}

	user := claims["user"].(map[string]interface{})
	if user["name"] == nil || user["role"] == nil {
		return models.User{}, fmt.Errorf("Authorization information invalid.")
	}

	return models.User{
		Name: user["name"].(string),
		Role: int64(user["role"].(float64)),
	}, nil
}
