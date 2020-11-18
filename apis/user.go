package apis

import (
	"../models"
	"../utils"
	"fmt"
	"net/http"
)

// UserAPI is handler for /user
func UserAPI(w http.ResponseWriter, r *http.Request) {
	user := GetUserFromRequest(r)
	if user.Role == models.NONUSER {
		utils.RespondCustomError(w, http.StatusUnauthorized, "Only registered user can access.")
		return
	}

	utils.RespondJson(w, http.StatusOK, fmt.Sprintf("Hi `%v`..! Welcome to user API.", user.Name))
}
