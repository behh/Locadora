package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/sessions"
	"github.com/urfave/negroni"
)

//BasicAuth adiciona o Handler para autenticação simples
func BasicAuth() negroni.Handler {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		user, pass, ok := r.BasicAuth()
		if !ok || user != "usuario" || pass != "senha" { /*apenas para exemplo*/
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, `{"error": "Unauthorized"}`)
			return
		}
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		next(w, r)
	})
}

//SessionAuth adiciona o Handler para verificar se está autenticado
func SessionAuth(store *sessions.CookieStore) negroni.Handler {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

		path := strings.Split(r.URL.Path, "/")
		if (r.URL.Path) == "/login" || (r.URL.Path) == "/loginauth" || path[1] == "static" {
			next(w, r)
			return
		}

		session, err := store.Get(r, "session")
		if err != nil {
			http.Error(w, fmt.Sprintf("erro ao obter sessão: %s", err), http.StatusInternalServerError)
			return
		}
		_, ok := session.Values["user"]
		if !ok {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		next(w, r)
	})
}
