package apis

import (
	logger "datumbrain/my-project/log"
	"datumbrain/my-project/models"
	"datumbrain/my-project/utils"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/casbin/casbin"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
)

var TokenAuth *jwtauth.JWTAuth
var authEnforcer *casbin.Enforcer

var signKey = []byte(os.Getenv("JWT_SIGN_KEY"))
var verifyKey = []byte(os.Getenv("JWT_VERIFY_KEY"))

func init() {
	TokenAuth = jwtauth.New("HS256", signKey, nil)
	var err error
	authEnforcer, err = casbin.NewEnforcerSafe("./config/model.conf", "./config/policy.csv")
	if err != nil {
		log.Fatal(err)
	}

}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var body models.SignInRequest
	if err := utils.ParseJson(r, &body); err != nil {
		logger.Error.Println(err)
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

	logger.Info.Printf("User `%v` signed in.\n", user.Name)
}

func verifyUserInfo(user models.SignInRequest) (models.User, error) {
	switch {
	case user.Username == "admin" && user.Password == "admin":
		return models.User{
			Name: user.Username,
			Role: models.USER_ADMIN,
		}, nil
	case user.Username == "john" && user.Password == "doe":
		return models.User{
			Name: user.Username,
			Role: models.USER_REGULAR,
		}, nil
	default:
		return models.User{}, fmt.Errorf("username or password invalid")
	}
}

// JWT Auth

// AdminAuthenticator checks if request is from an admin user
func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		user, err := getUserFromRequest(r)
		if err != nil {
			logger.Warn.Println("Someone tried to access protected method without authentication.")
			utils.RespondCustomError(w, http.StatusUnauthorized, "Authorization information missing/invalid.")
			return
		}

		res, err := authEnforcer.EnforceSafe(user.Role, r.URL.Path, r.Method)
		if err != nil {
			writeError(http.StatusInternalServerError, "ERROR", w, err)
			return
		}
		if res == false {
			writeError(http.StatusForbidden, "FORBIDDEN", w, errors.New("user does not exist"))
			return			
		}

		next.ServeHTTP(w, r)

	})
}

func writeError(status int, message string, w http.ResponseWriter, err error) {
	log.Print("ERROR: ", err.Error())
	w.WriteHeader(status)
	w.Write([]byte(message))
}

func getUserFromRequest(r *http.Request) (models.User, error) {
	token, claims, err := jwtauth.FromContext(r.Context())

	if err != nil || token == nil || !token.Valid {
		logger.Error.Printf("%v\n", err)
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
