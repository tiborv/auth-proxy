package main

import (
	"log"
	"net/http"
	"time"

	"github.com/tiborv/prxy/db"
	"github.com/tiborv/prxy/routes"
)

func main() {
	db.Connect()

	s := &http.Server{
		Addr:           ":3001",
		Handler:        routes.GetRootMux(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
