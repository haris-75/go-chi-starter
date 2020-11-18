package apis

import (
	"../constants"
	"../log"
	"../models"
	"../utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"net/http"
)

var TokenAuth *jwtauth.JWTAuth

func init() {
	TokenAuth = jwtauth.New("HS256", []byte(constants.SignKey), nil)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var body models.SignInRequest
	if err := utils.ParseJson(r, &body); err != nil {
		log.Error.Println(err)
		utils.RespondError(w, http.StatusBadRequest)
		return
	}

	user := verifyUserInfo(body)
	if user.Role == models.NONUSER {
		utils.RespondCustomError(w, http.StatusNotFound, "Invalid username or password.")
		return
	}
	_, tokenString, _ := TokenAuth.Encode(jwt.MapClaims{"user": user})
	w.Header().Set("JWT-Token", tokenString)
	log.Info.Printf("User `%v` signed in.\n", user.Name)
	utils.RespondJson(w, http.StatusOK, user)
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
