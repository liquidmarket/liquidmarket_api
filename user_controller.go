package liquidmarket

import (
	"net/http"
)

func (a *App) getUsers(w http.ResponseWriter, r *http.Request) {
    users, err := GetUsers(a.DB)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }
    respondWithJSON(w, http.StatusOK, users)
}