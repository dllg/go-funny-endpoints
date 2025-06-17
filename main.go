package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dllg/go-funny-endpoints/config"
	"github.com/dllg/go-funny-endpoints/router"
)

func main() {
	config.Init()                         // Initialize configuration settings
	port := config.Get(config.ServerPort) // Get the server port from config
	addr := ":" + port

	mux := router.Setup()

	// attach an ErrorLog so you’ll see Listen/Serve errors
	srv := &http.Server{
		Addr:     addr,
		Handler:  mux,
		ErrorLog: log.New(os.Stderr, "http-server: ", log.LstdFlags),
	}

	log.Printf("listening on %s…\n", addr)
	log.Fatal(srv.ListenAndServe())
}
