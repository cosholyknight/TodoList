package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"time"
)

type application struct {
	config config
}

type config struct {
	addr string
}

// use chi package to routing
func (app *application) mount() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	return r
}

func (app *application) run(mux *http.ServeMux) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("server has started at %s", app.config.addr)
	return srv.ListenAndServe()
}
