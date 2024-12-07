package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type config struct {
	port int
}

type application struct {
	config config
	logger *log.Logger
}

const version = "1.23"

func main() {
	var config config

	flag.IntVar(&config.port, "port", 8080, "Application starting port")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: config,
		logger: logger,
	}

	srv := http.Server{
		Addr:         fmt.Sprintf(":%d", config.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err := srv.ListenAndServe()
	log.Fatal(err)
}
