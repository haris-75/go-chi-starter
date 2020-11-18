package apis

import (
	"datumbrain/my-project/log"
	"datumbrain/my-project/models"
	"datumbrain/my-project/utils"
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
func AdminAuthenticator(next http.Handler) http.Handler {
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

// UserAuthenticator checks if request is from an valid user
func UserAuthenticator(next http.Handler) http.Handler {
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
