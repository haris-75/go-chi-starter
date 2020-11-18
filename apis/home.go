package apis

import (
	"datumbrain/my-project/utils"
	"net/http"
)

// HomeAPI is handler for /
func HomeAPI(w http.ResponseWriter, r *http.Request) {
	utils.RespondJson(w, http.StatusOK, "API seems fine.")
}
