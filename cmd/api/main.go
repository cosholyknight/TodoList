package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	// use gotdotenv package to load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// config takes addr from .env
	cfg := config{
		addr: os.Getenv("ADDR"),
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
