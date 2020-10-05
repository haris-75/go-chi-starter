package apis

import (
	"../models"
	"fmt"
	"github.com/go-chi/jwtauth"
	"net/http"
)

// HomeAPI is handler for /
func HomeAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API seems fine.")
}

// AdminAPI is handler for /admin
func AdminAPI(w http.ResponseWriter, r *http.Request) {
	user := getUser(r)
	if user.Role != models.USER_ADMIN {
		http.Error(w, "Only admin can access.", http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "Hi `%v`..! Welcome to admin API.", user.Name)
}

// UserAPI is handler for /user
func UserAPI(w http.ResponseWriter, r *http.Request) {
	user := getUser(r)
	if user.Role == models.USER_NONUSER {
		http.Error(w, "Only registered user can access.", http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "Hi `%v`..! Welcome to user API.", user.Name)
}

func getUser(r *http.Request) models.User {
	_, claims, _ := jwtauth.FromContext(r.Context())
	user := claims["user"].(map[string]interface{})
	if user["name"] == nil || user["role"] == nil {
		return models.User{"", models.USER_NONUSER}
	}
	return models.User{user["name"].(string), int64(user["role"].(float64))}
}
