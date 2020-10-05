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
	fmt.Fprintf(w, "protected area. hi %v", claims["user_id"])
}
