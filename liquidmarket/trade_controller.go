package liquidmarket

import (
	"encoding/json"
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

func (a *App) submitTrade(w http.ResponseWriter, r *http.Request) {
	var tt TradeToken
	user, err := insecureGetUserFromJWT(r)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Problem with the JWT")
		return
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tt); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	trade, err := tt.trade(a.DB, *user)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, trade)
}
