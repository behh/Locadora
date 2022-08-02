package handler

import (
	"fmt"
	"net/http"
	"os"
	"text/template"

	"github.com/behh/locadora/usecase/logger"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/urfave/negroni"
)

func MakeAuthHandlers(r *mux.Router, n negroni.Negroni, store *sessions.CookieStore, ls logger.UseCase) {
	r.Handle("/login", n.With(
		negroni.Wrap(login(ls)),
	)).Methods("GET").Name("login")

	r.Handle("/loginauth", n.With(
		negroni.Wrap(loginAuth(store, ls)),
	)).Methods("POST").Name("loginauth")

	r.Handle("/logout", n.With(
		negroni.Wrap(logout(store, ls)),
	)).Methods("GET").Name("logout")
}

func login(ls logger.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		pathTemplates := os.Getenv("PATH_TEMPLATES")
		ts, err := template.ParseFiles(
			pathTemplates+"/header.html",
			pathTemplates+"/login.html",
			pathTemplates+"/footer.html")
		if err != nil {
			ls.LogError(err)
			http.Error(w, fmt.Sprintf("error parsing: %s", err), http.StatusInternalServerError)
			return
		}

		data := struct {
			Title  string
			Result string
		}{
			Title:  "Login",
			Result: "",
		}
		err = ts.Lookup("login.html").ExecuteTemplate(w, "login", data)
		if err != nil {
			ls.LogError(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}

func loginAuth(store *sessions.CookieStore, ls logger.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		pathTemplates := os.Getenv("PATH_TEMPLATES")
		ts, err := template.ParseFiles(
			pathTemplates+"/header.html",
			pathTemplates+"/login.html",
			pathTemplates+"/index.html",
			pathTemplates+"/navbar.html",
			pathTemplates+"/footer.html")
		if err != nil {
			ls.LogError(err)
			http.Error(w, fmt.Sprintf("error parsing: %s", err), http.StatusInternalServerError)
			return
		}
		r.ParseForm()
		user := r.FormValue("user")
		pass := r.FormValue("pass")

		data := struct {
			Title  string
			Result string
		}{
			Title:  "Login",
			Result: "Usuário Logado com Sucesso",
		}
		if user != "usuario" || pass != "senha" { /*apenas para exemplo*/
			data.Result = "Usuário e Senha Inválidos"
			err = ts.Lookup("login.html").ExecuteTemplate(w, "login", data)
			if err != nil {
				ls.LogError(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		session, err := store.Get(r, "session")
		if err != nil {
			ls.LogError(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		session.Values["user"] = user
		session.Save(r, w)
		err = ts.Lookup("index.html").ExecuteTemplate(w, "index", data)
		if err != nil {
			ls.LogError(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}

func logout(store *sessions.CookieStore, ls logger.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		session, err := store.Get(r, "session")
		if err != nil {
			ls.LogError(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		session.Options.MaxAge = -1
		session.Save(r, w)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	})
}
