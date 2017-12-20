package liquidmarket

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
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

func (a *App) getAccountsOrCreateUser(w http.ResponseWriter, r *http.Request) {
	user, err := insecureGetUserFromJWT(r)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}
	accounts, err := user.getAccountsOrCreate(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, accounts)
}

func (a *App) getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, err := insecureGetUserFromJWT(r)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}
	if user.GoogleID == vars["google_id"] {
		u := User{GoogleID: vars["google_id"]}
		if err := u.GetUser(a.DB); err != nil {
			switch err {
			case sql.ErrNoRows:
				respondWithError(w, http.StatusNotFound, "Account not found")
			default:
				respondWithError(w, http.StatusInternalServerError, err.Error())
			}
			return
		}
		respondWithJSON(w, http.StatusOK, u)
	} else {
		respondWithError(w, http.StatusForbidden, "user id doesn't match JWT")
	}
}
