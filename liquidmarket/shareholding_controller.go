package liquidmarket

import (
	"github.com/satori/go.uuid"
	"net/http"
    "github.com/gorilla/mux"
)

func (a *App) getShareholdings(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid, err := uuid.FromString(vars["uuid"])
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid Account ID")
        return
    }
    users, err := GetShareHoldings(a.DB, guid)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }
    respondWithJSON(w, http.StatusOK, users)
}