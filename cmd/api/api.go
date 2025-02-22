package main

import (
	"TodoList/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"time"
)

type application struct {
	config config
	store  store.Storage
}
type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

// use chi package to routing
func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	// logger for http request
	r.Use(middleware.Logger)

	// recover from panic
	r.Use(middleware.Recoverer)
	return r
}

func (app *application) run(mux http.Handler) error {
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
