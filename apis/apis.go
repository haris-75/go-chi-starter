package apis

import (
	"../models"
	"fmt"
	"github.com/go-chi/jwtauth"
	"net/http"
)

// HomeAPI is handler for /
func HomeAPI(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, "API seems fine.")
}

// AdminAPI is handler for /admin
func AdminAPI(w http.ResponseWriter, r *http.Request) {
	user := getUserFromRequest(r)
	if user.Role != models.USER_ADMIN {
		RespondError(w, http.StatusUnauthorized, "Only admin can access.")
		return
	}

	RespondJSON(w, http.StatusOK, fmt.Sprintf("Hi `%v`..! Welcome to admin API.", user.Name))
}

// UserAPI is handler for /user
func UserAPI(w http.ResponseWriter, r *http.Request) {
	user := getUserFromRequest(r)
	if user.Role == models.NONUSER {
		RespondError(w, http.StatusUnauthorized, "Only registered user can access.")
		return
	}

	RespondJSON(w, http.StatusOK, fmt.Sprintf("Hi `%v`..! Welcome to user API.", user.Name))
}

func getUserFromRequest(r *http.Request) models.User {
	_, claims, _ := jwtauth.FromContext(r.Context())
	user := claims["user"].(map[string]interface{})
	if user["name"] == nil || user["role"] == nil {
		return models.User{"", models.NONUSER}
	}
	return models.User{user["name"].(string), int64(user["role"].(float64))}
}
