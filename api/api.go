package api

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/behh/locadora/api/handler"
	"github.com/behh/locadora/api/middleware"
	"github.com/behh/locadora/config"
	"github.com/behh/locadora/infrastructure/repository"
	"github.com/behh/locadora/usecase/carro"
	"github.com/behh/locadora/usecase/logger"
	whandler "github.com/behh/locadora/web/handler"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/urfave/negroni"
)

func InitAPI() {

	ls := logger.NewService()

	db, err := repository.InitDBMSSQL(config.DB_HOST, config.DB_PORT, config.DB_USER, config.DB_PASSWORD, config.DB_DATABASE)
	if err != nil {
		ls.Logger.Error.Fatal(err)
	}
	defer db.Close()

	carroRepo := repository.NewCarroMSSQL(db)
	carroService := carro.NewService(carroRepo)

	r := mux.NewRouter()

	//Inicio API
	n := negroni.New()
	n.Use(middleware.APILogger())
	n.Use(middleware.Cors())
	n.Use(middleware.BasicAuth())
	n.Use(middleware.ApplicationJSON())

	//carro
	handler.MakeCarroHandlers(r, *n, carroService, ls)

	http.Handle("/", r)
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	//Fim API

	//Inicio WEB
	store := sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	os.Setenv("PATH_TEMPLATES", config.PATH_TEMPLATES)

	n2 := negroni.New()
	n2.Use(middleware.APILogger())
	n2.Use(middleware.SessionAuth(store))

	whandler.MakeCarroHandlers(r, *n2, carroService, ls)
	whandler.MakeAuthHandlers(r, *n2, store, ls)

	//static files
	fileServer := http.FileServer(http.Dir(config.PATH_STATIC))
	r.PathPrefix("/static/").Handler(n2.With(
		negroni.Wrap(http.StripPrefix("/static/", fileServer)),
	)).Methods("GET", "OPTIONS")
	//Fim WEB

	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + strconv.Itoa(config.API_PORT),
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     ls.Logger.Error,
	}
	ls.Logger.Info.Printf("API da Locadora Rodando na Porta: %d\n", config.API_PORT)
	if config.TLS {
		err = srv.ListenAndServeTLS(config.PATH_CERT, config.PATH_KEY)
		if err != nil {
			ls.Logger.Error.Fatal(err.Error())
		}
	} else {
		err = srv.ListenAndServe()
		if err != nil {
			ls.Logger.Error.Fatal(err.Error())
		}
	}

}
