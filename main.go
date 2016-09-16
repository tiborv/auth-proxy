package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/tiborv/prxy/models"
	"github.com/tiborv/prxy/routes"
)

func main() {
	models.Connect()

	s := &http.Server{
		Addr:           ":" + os.Getenv("PORT"),
		Handler:        routes.GetRootMux(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
