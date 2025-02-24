package main

import (
	"TodoList/internal/store"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"strconv"
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

	// routes for lists table
	r.Route("/lists", func(r chi.Router) {
		r.Post("/", app.createList)
		r.Get("/", app.GetTasksFromList)
		r.Delete("/{listID}", app.deleteList)
	})
	r.Route("/tasks", func(r chi.Router) {
		r.Post("/", app.createTask)
		r.Delete("/{taskID}", app.DeleteTask)
	})
	return r
}

// run server
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

// List API
// Create a list
func (app *application) createList(w http.ResponseWriter, r *http.Request) {
	var listBody store.List
	err := json.NewDecoder(r.Body).Decode(&listBody)
	if err != nil {
		http.Error(w, "Failed to parse body", http.StatusBadRequest)
	}

	err = app.store.Lists.Create(r.Context(), &listBody)
	if err != nil {
		http.Error(w, "Failed to create list", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(listBody)
}

// Delete a list
func (app *application) deleteList(w http.ResponseWriter, r *http.Request) {
	listIDAsString := chi.URLParam(r, "listID")
	listID, err := strconv.Atoi(listIDAsString)
	if err != nil {
		http.Error(w, "Failed to convert listID", http.StatusBadRequest)
	}

	err = app.store.Lists.Delete(r.Context(), int64(listID))
	if err != nil {
		http.Error(w, "Failed to delete list", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

// Task api
// Create a task
func (app *application) createTask(w http.ResponseWriter, r *http.Request) {
	var taskBody store.Task
	err := json.NewDecoder(r.Body).Decode(&taskBody)
	if err != nil {
		http.Error(w, "Failed to parse task body", http.StatusBadRequest)
	}
	err = app.store.Tasks.Create(r.Context(), &taskBody)
	if err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(taskBody)
}

func (app *application) DeleteTask(w http.ResponseWriter, r *http.Request) {
	taskIDAsString := chi.URLParam(r, "taskID")
	taskID, err := strconv.Atoi(taskIDAsString)
	if err != nil {
		http.Error(w, "Failed to parse taskID", http.StatusBadRequest)
	}
	err = app.store.Tasks.Delete(r.Context(), int64(taskID))
	if err != nil {
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
