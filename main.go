package main

import (
    "net/http"
	"os"
	"log"

	"github.com/dllg/go-funny-endpoints/router"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}
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
