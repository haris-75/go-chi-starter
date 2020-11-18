package apis

import (
	"../models"
	"../utils"
	"fmt"
	"net/http"
)

// AdminAPI is handler for /admin
func AdminAPI(w http.ResponseWriter, r *http.Request) {
	user := GetUserFromRequest(r)
	if user.Role != models.USER_ADMIN {
		utils.RespondCustomError(w, http.StatusUnauthorized, "Only admin can access")
		return
	}

	utils.RespondJson(w, http.StatusOK, fmt.Sprintf("Hi `%v`..! Welcome to admin API.", user.Name))
}
