package liquidmarket

import (
	"encoding/json"
	"net/http"
)

func (a *App) submitOffer(w http.ResponseWriter, r *http.Request) {
	var offer Offer
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&offer); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	if err := offer.submitOffer(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, offer)
}
