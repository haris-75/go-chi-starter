package apis

import (
	"../constants"
	"../models"
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
	ParseJSON(r, &body)

	user := verifyUserInfo(body)
	if user.Role == models.NONUSER {
		RespondError(w, http.StatusNotFound, "Invalid username or password.")
		return
	}
	_, tokenString, _ := TokenAuth.Encode(jwt.MapClaims{"user": user})
	w.Header().Set("JWT-Token", tokenString)
	log.Printf("[RUN]\tUser `%v` signed in.\n", user.Name)
	RespondJSON(w, http.StatusOK, user)
}

func verifyUserInfo(user models.SignInRequest) models.User {
	switch {
	case user.Username == "admin" && user.Password == "admin":
		return models.User{user.Username, models.USER_ADMIN}
	case user.Username == "faizan" && user.Password == "faizan":
		return models.User{user.Username, models.USER_REGULAR}
	default:
		return models.User{user.Username, models.NONUSER}
	}
}

// GetUserFromRequest gets user information from JWT-Token
func GetUserFromRequest(r *http.Request) models.User {
	_, claims, _ := jwtauth.FromContext(r.Context())
	user := claims["user"].(map[string]interface{})
	if user["name"] == nil || user["role"] == nil {
		return models.User{"", models.NONUSER}
	}
	return models.User{user["name"].(string), int64(user["role"].(float64))}
}
