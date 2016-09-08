package main

import (
	"log"
	"net/http"
	"time"

	db "./db"
	routes "./routes"
)

func main() {
	db.Connect()

	s := &http.Server{
		Addr:           ":3001",
		Handler:        routes.RootHandler(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
