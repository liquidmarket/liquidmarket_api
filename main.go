package liquidmarket

import (
	"net/http"
)

func init() {
	a := App{} 

	a.Initialize("root", "podsaveamerica", "enhanced-emblem-188503:australia-southeast1:liquidmarket", "liquidmarket")
	
	a.Run()
}

func handler(w http.ResponseWriter, r *http.Request) {
	user, err := insecureGetUserFromJWT(r)

	if err == nil {
		respondWithJSON(w, http.StatusOK, user)
	} else {
		respondWithError(w, http.StatusUnauthorized, err.Error())
	}
}
