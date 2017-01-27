package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/tiborv/auth-proxy/models"
	"github.com/tiborv/auth-proxy/routes"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}
	models.Connect(true)

	fmt.Println(port)
	s := &http.Server{
		Addr:           ":" + port,
		Handler:        routes.GetRootMux(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
