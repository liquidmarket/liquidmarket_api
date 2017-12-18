package liquidmarket

import (
	"net/http"
)

func (a *App) getPrices(w http.ResponseWriter, r *http.Request) {
    users, err := GetPrices(a.DB)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }
    respondWithJSON(w, http.StatusOK, users)
}