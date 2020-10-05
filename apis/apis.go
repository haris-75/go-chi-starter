package apis

import (
	"fmt"
	"github.com/go-chi/jwtauth"
	"net/http"
)

// Home is handler for /
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API seems fine.")
}

// Test is handler for /test
func Test(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	user := claims["user"].(map[string]interface{})

	fmt.Fprintf(w, "Hi `%v`..! Welcome to your private function.", user["username"])
}
