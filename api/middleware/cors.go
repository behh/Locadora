package middleware

import (
	"net/http"

	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

//Cors adiciona os headers para suportar o CORS nos navegadores
func Cors() negroni.Handler {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowedHeaders:   []string{"Content-Type", "api_key", "Authorization"},
		AllowCredentials: true,
	})
}
