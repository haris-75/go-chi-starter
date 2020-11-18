package apis

import (
	"../log"
	"../models"
	"../utils"
	"fmt"
	"net/http"
)

// AdminAPI is handler for /admin
func AdminAPI(w http.ResponseWriter, r *http.Request) {
	user := GetUserFromRequest(r)
	if user.Role != models.USER_ADMIN {
		log.Warn.Printf("`%v` tried to access protected method.\n", user.Name)
		utils.RespondCustomError(w, http.StatusUnauthorized, "Only admin can access")
		return
	}

	utils.RespondJson(w, http.StatusOK, fmt.Sprintf("Hi `%v`..! Welcome to admin API.", user.Name))
}
