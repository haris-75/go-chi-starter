package apis

import (
	"../utils"
	"fmt"
	"net/http"
)

// UserAPI is handler for /user
func UserAPI(w http.ResponseWriter, r *http.Request) {
	utils.RespondJson(w, http.StatusOK, fmt.Sprintf("Welcome to user api."))
}
