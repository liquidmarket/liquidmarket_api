package liquidmarket

import (
	"net/http"
)

func (a *App) updateUser(w http.ResponseWriter, r *http.Request) {
	user, err := insecureGetUserFromJWT(r)
	if err != nil {
        respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}
    if err := user.UpdateUser(a.DB); err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }
    respondWithJSON(w, http.StatusOK, user)
}