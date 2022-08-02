package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/behh/locadora/api/presenter"
	"github.com/behh/locadora/entity"
	"github.com/behh/locadora/usecase/carro"
	"github.com/behh/locadora/usecase/logger"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func MakeCarroHandlers(r *mux.Router, n negroni.Negroni, service carro.UseCase, ls logger.UseCase) {
	r.Handle("/v1/carro", n.With(
		negroni.Wrap(listCarros(service, ls)),
	)).Methods("GET", "OPTIONS").Name("listCarros")

	r.Handle("/v1/carro", n.With(
		negroni.Wrap(createCarro(service, ls)),
	)).Methods("POST", "OPTIONS").Name("createCarro")

	r.Handle("/v1/carro/{id}", n.With(
		negroni.Wrap(getCarro(service, ls)),
	)).Methods("GET", "OPTIONS").Name("getCarro")

	r.Handle("/v1/carro/{id}", n.With(
		negroni.Wrap(deleteCarro(service, ls)),
	)).Methods("DELETE", "OPTIONS").Name("deleteCarro")
}

func listCarros(service carro.UseCase, ls logger.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "erro na busca de carros"
		var data []*entity.Carro
		var err error

		placa := r.URL.Query().Get("placa")
		switch {
		case len(placa) == 0:
			data, err = service.ListCarros()
		default:
			data, err = service.SearchCarros(placa)
		}

		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			ls.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			errorMessage = "pacote vazio"
			ls.LogInfo(errorMessage)
			w.WriteHeader(http.StatusNoContent)
			w.Write([]byte(errorMessage))
			return
		}
		var toJ []*presenter.Carro
		for _, d := range data {
			toJ = append(toJ, &presenter.Carro{
				ID:          d.ID,
				Placa:       d.Placa,
				Modelo:      d.Modelo,
				Cor:         d.Cor,
				Renavan:     d.Renavan,
				Hodometro:   d.Hodometro,
				DataCriacao: d.DataCriacao,
			})
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			ls.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func createCarro(service carro.UseCase, ls logger.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "erro na inserção de carros"
		var input struct {
			Placa     string `json:"placa"`
			Modelo    string `json:"modelo"`
			Cor       string `json:"cor"`
			Renavan   int    `json:"renavan"`
			Hodometro int    `json:"hodometro"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			ls.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		id, err := service.CreateCarro(input.Placa, input.Modelo, input.Cor, input.Renavan, input.Hodometro)
		if err != nil {
			ls.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		toJ := &presenter.Carro{
			ID:        id,
			Placa:     input.Placa,
			Modelo:    input.Modelo,
			Cor:       input.Cor,
			Renavan:   input.Renavan,
			Hodometro: input.Hodometro,
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func getCarro(service carro.UseCase, ls logger.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "erro na busca do carro"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			ls.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		data, err := service.GetCarro(id)
		if err != nil && err != entity.ErrNotFound {
			ls.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNoContent)
			w.Write([]byte(errorMessage))
			return
		}
		toJ := &presenter.Carro{
			ID:          data.ID,
			Placa:       data.Placa,
			Modelo:      data.Modelo,
			Cor:         data.Cor,
			Renavan:     data.Renavan,
			Hodometro:   data.Hodometro,
			DataCriacao: data.DataCriacao,
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			ls.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func deleteCarro(service carro.UseCase, ls logger.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "erro na remoção do carro"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			ls.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		err = service.DeleteCarro(id)
		if err != nil {
			ls.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}
