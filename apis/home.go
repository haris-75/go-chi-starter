package apis

import (
	"net/http"
)

// HomeAPI is handler for /
func HomeAPI(w http.ResponseWriter, r *http.Request) {
	RespondJSON(w, http.StatusOK, "API seems fine.")
}
