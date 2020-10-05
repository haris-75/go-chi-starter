package apis

import (
	"../models"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"log"
	"net/http"
)

var TokenAuth *jwtauth.JWTAuth

func init() {
	TokenAuth = jwtauth.New("HS256", []byte("secret_key_here"), nil)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var user models.SignInRequest

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Printf("[ERROR]\t%v", err)
		http.Error(w, "Unable to parse json.", http.StatusBadRequest)
		return
	}

	if isVerified(user) {
		_, tokenString, _ := TokenAuth.Encode(jwt.MapClaims{"user": user})
		w.Header().Set("JWT-Token", tokenString)
	} else {
		http.Error(w, "inavlid username/password.", http.StatusNotFound)
	}
}

func isVerified(user models.SignInRequest) bool {
	return user.Username == "admin" && user.Password == "admin"
}
