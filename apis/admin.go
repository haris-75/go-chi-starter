package apis

import (
	"../utils"
	"fmt"
	"net/http"
)

// AdminAPI is handler for /admin
func AdminAPI(w http.ResponseWriter, r *http.Request) {
	utils.RespondJson(w, http.StatusOK, fmt.Sprintf("Welcome to admin api."))
}
