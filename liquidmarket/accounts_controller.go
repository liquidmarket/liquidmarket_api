package liquidmarket

import (
    "net/http"
    "github.com/gorilla/mux"
    "github.com/satori/go.uuid"
    "database/sql"
    "fmt"
    "encoding/json"
)

func (a *App) getAccount(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    x, err := uuid.FromString(vars["uuid"])
    if err != nil {
        respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid account ID: %s", err.Error()))
        return
    }
	acc := Account{ID: x}
    if err := acc.GetAccount(a.DB); err != nil {
        switch err {
        case sql.ErrNoRows:
            respondWithError(w, http.StatusNotFound, "Account not found")
        default:
            respondWithError(w, http.StatusInternalServerError, err.Error())
        }
        return
    }
    respondWithJSON(w, http.StatusOK, acc)
}

func (a *App) getAccounts(w http.ResponseWriter, r *http.Request) {
    user, err := insecureGetUserFromJWT(r)
    if err != nil {
        respondWithError(w, http.StatusUnauthorized, "Problem with the JWT")
        return
    }
    accounts, err := GetAccounts(a.DB, user.GoogleID)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }
    respondWithJSON(w, http.StatusOK, accounts)
}

func (a *App) createAccount(w http.ResponseWriter, r *http.Request) {
	var acc Account
    user, err := insecureGetUserFromJWT(r)
    if err != nil {
        respondWithError(w, http.StatusUnauthorized, "Problem with the JWT")
        return
    }
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&acc); err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }
	defer r.Body.Close()
    if err := acc.CreateAccount(a.DB, user.GoogleID); err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }
    respondWithJSON(w, http.StatusCreated, acc)
}

func (a *App) updateAccount(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    var acc Account
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&acc); err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
        return
    }
    defer r.Body.Close()
    guid, err := uuid.FromString(vars["uuid"])
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid Account ID")
        return
    }
    if err := acc.UpdateAccount(a.DB, guid); err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }
    respondWithJSON(w, http.StatusOK, acc)
}

func (a *App) deleteAccount(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    guid, err := uuid.FromString(vars["uuid"])
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid Account ID")
        return
    }
    if err := DeleteAccount(a.DB, guid); err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }
    respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}