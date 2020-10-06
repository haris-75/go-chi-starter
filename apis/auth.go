package apis

import (
	"../constants"
	"../models"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"log"
	"net/http"
)

var TokenAuth *jwtauth.JWTAuth

func init() {
	TokenAuth = jwtauth.New("HS256", []byte(constants.SignKey), nil)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var body models.SignInRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		log.Printf("[ERROR]\t%v", err)
		http.Error(w, "Unable to parse json.", http.StatusBadRequest)
		return
	}

	user := getUserInfo(body)
	if user.Role == models.NONUSER {
		http.Error(w, "inavlid username/password.", http.StatusNotFound)
		return
	}
	_, tokenString, _ := TokenAuth.Encode(jwt.MapClaims{"user": user})
	w.Header().Set("JWT-Token", tokenString)
	log.Printf("[RUN]\tUser `%v` signed in.\n", user.Name)
}

func getUserInfo(user models.SignInRequest) models.User {
	switch {
	case user.Username == "admin" && user.Password == "admin":
		return models.User{user.Username, models.USER_ADMIN}
	case user.Username == "faizan" && user.Password == "faizan":
		return models.User{user.Username, models.USER_REGULAR}
	default:
		return models.User{user.Username, models.NONUSER}
	}
}
