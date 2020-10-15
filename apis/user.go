package apis

import (
	"../models"
	"fmt"
	"net/http"
)

// UserAPI is handler for /user
func UserAPI(w http.ResponseWriter, r *http.Request) {
	user := GetUserFromRequest(r)
	if user.Role == models.NONUSER {
		RespondError(w, http.StatusUnauthorized, "Only registered user can access.")
		return
	}

	RespondJSON(w, http.StatusOK, fmt.Sprintf("Hi `%v`..! Welcome to user API.", user.Name))
}
