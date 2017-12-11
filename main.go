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
	user := User{ 
		GoogleID: "111307193244795741281",
		FirstName: "Thomas",
		LastName: "Horrobin",
		Email: "thomasroberthorrobin@gmail.com",
	}

	respondWithJSON(w, http.StatusOK, user)
}
