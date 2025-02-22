package main

import (
	"TodoList/internal/db"
	"TodoList/internal/store"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strconv"
)

func main() {
	// use gotdotenv package to load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	maxIdleConns, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNS"))
	if err != nil {
		log.Fatal("error converting maxIdleConns")
	}

	maxOpenConns, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNS"))
	if err != nil {
		log.Fatal("error converting maxIdleConns")
	}

	// config takes addr from .env
	cfg := config{
		addr: os.Getenv("ADDR"),
		db: dbConfig{
			addr:         os.Getenv("DB_ADDR"),
			maxOpenConns: maxOpenConns,
			maxIdleConns: maxIdleConns,
			maxIdleTime:  os.Getenv("DB_MAX_IDLE_TIME"),
		},
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("db connection pool established")

	storage := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  storage,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
