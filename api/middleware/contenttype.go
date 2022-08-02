package middleware

import (
	"net/http"

	"github.com/urfave/negroni"
)

//ApplicationJSON é o Middleware que adiciona o Content-Type
func ApplicationJSON() negroni.Handler {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	})
}
