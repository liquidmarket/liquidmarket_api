package liquidmarket

import (
	"net/http"
)

func (a *App) getTrades(w http.ResponseWriter, r *http.Request) {
	user, err := insecureGetUserFromJWT(r)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Problem with the JWT")
		return
	}
	trades, err := getTrades(a.DB, user.GoogleID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, trades)
}
