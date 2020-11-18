package apis

import (
	"../log"
	"../models"
	"../utils"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"net/http"
	"os"
)

var TokenAuth *jwtauth.JWTAuth

var signKey = []byte(os.Getenv("JWT_SIGN_KEY"))
var verifyKey = []byte(os.Getenv("JWT_VERIFY_KEY"))

func init() {
	TokenAuth = jwtauth.New("HS256", signKey, nil)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var body models.SignInRequest
	if err := utils.ParseJson(r, &body); err != nil {
		log.Error.Println(err)
		utils.RespondError(w, http.StatusBadRequest)
		return
	}

	user, err := verifyUserInfo(body)
	if err != nil {
		utils.RespondCustomError(w, http.StatusNotFound, err.Error())
		return
	}

	_, tokenString, _ := TokenAuth.Encode(jwt.MapClaims{"user": user})

	w.Header().Set("JWT-Token", tokenString)
	utils.RespondJson(w, http.StatusOK, user)

	log.Info.Printf("User `%v` signed in.\n", user.Name)
}

func verifyUserInfo(user models.SignInRequest) (models.User, error) {
	switch {
	case user.Username == "admin" && user.Password == "admin":
		return models.User{
			Name: user.Username,
			Role: models.USER_ADMIN,
		}, nil
	case user.Username == "faizan" && user.Password == "faizan":
		return models.User{
			Name: user.Username,
			Role: models.USER_REGULAR,
		}, nil
	default:
		return models.User{}, fmt.Errorf("username or password invalid")
	}
}
