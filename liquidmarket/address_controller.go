package liquidmarket

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
)

func (a *App) getAddresses(w http.ResponseWriter, r *http.Request) {
	user, err := insecureGetUserFromJWT(r)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Problem with the JWT")
		return
	}
	addresses, err := GetAddresss(a.DB, user.GoogleID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, addresses)
}

func (a *App) createAddress(w http.ResponseWriter, r *http.Request) {
	var address Address
	user, err := insecureGetUserFromJWT(r)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Problem with the JWT")
		return
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&address); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	address.GoogleID = user.GoogleID
	defer r.Body.Close()
	if err := address.CreateAddress(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, address)
}

func (a *App) updateAddress(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var address Address
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&address); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	guid, err := uuid.FromString(vars["uuid"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Address ID")
		return
	}
	if err := UpdateAddress(a.DB, guid, address.Address); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, address)
}

func (a *App) deleteAddress(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	guid, err := uuid.FromString(vars["uuid"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Address ID")
		return
	}
	if err := DeleteAddress(a.DB, guid); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
