package liquidmarket

import (
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	user := User{ 
		ID: "111307193244795741281",
		FirstName: "Thomas",
		LastName: "Horrobin",
		Email: "thomasroberthorrobin@gmail.com",
	}

	respondWithJSON(w, http.StatusOK, user)
}
