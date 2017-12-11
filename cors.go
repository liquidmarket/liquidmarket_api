package liquidmarket

import (
	"net/http"
	"github.com/rs/cors"
)

func corsWrapper(h http.Handler) http.Handler {
	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowedHeaders: []string {"Authorization", "X-Requested-With", "Content-Type"},
		AllowedMethods: []string {"OPTIONS", "GET", "POST", "DELETE", "PUT"},
	})
	
	return handler.Handler(h)
}