package handler

import (
	"fmt"
	"net/http"
	"os"
	"text/template"

	"github.com/behh/locadora/entity"
	"github.com/behh/locadora/usecase/carro"
	"github.com/behh/locadora/usecase/logger"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func MakeCarroHandlers(r *mux.Router, n negroni.Negroni, service carro.UseCase, ls logger.UseCase) {
	r.Handle("/carros", n.With(
		negroni.Wrap(listCarros(service, ls)),
	)).Methods("GET", "OPTIONS").Name("listCarros")

}

func listCarros(service carro.UseCase, ls logger.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		pathTemplates := os.Getenv("PATH_TEMPLATES")
		ts, err := template.ParseFiles(
			pathTemplates+"/header.html",
			pathTemplates+"/navbar.html",
			pathTemplates+"/carros.html",
			pathTemplates+"/footer.html")
		if err != nil {
			ls.LogError(err)
			http.Error(w, fmt.Sprintf("error parsing: %s", err), http.StatusInternalServerError)
			return
		}

		placa := r.FormValue("placa")
		var carros []*entity.Carro

		if len(placa) > 0 {
			listaCarros, err := service.SearchCarros(placa)
			if err != nil {
				ls.LogError(err)
				http.Error(w, fmt.Sprintf("error parsing: %s", err), http.StatusInternalServerError)
				return
			}
			carros = append(carros, listaCarros...)
		} else {
			listaCarros, err := service.ListCarros()
			if err != nil {
				ls.LogError(err)
				http.Error(w, fmt.Sprintf("error parsing: %s", err), http.StatusInternalServerError)
				return
			}
			carros = append(carros, listaCarros...)
		}
		data := struct {
			Title  string
			Carros []*entity.Carro
		}{
			Title:  "Carros",
			Carros: carros,
		}
		err = ts.Lookup("carros.html").ExecuteTemplate(w, "carros", data)
		if err != nil {
			ls.LogError(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
