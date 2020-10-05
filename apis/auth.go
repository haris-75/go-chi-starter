package apis

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"net/http"
)

var TokenAuth *jwtauth.JWTAuth

func init() {
	TokenAuth = jwtauth.New("HS256", []byte("secret_key_here"), nil)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	_, tokenString, _ := TokenAuth.Encode(jwt.MapClaims{"user_id": 123})
	w.Header().Set("JWT-Token", tokenString)
}
