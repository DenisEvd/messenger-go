package main

import (
	"log"
	"messenger-go/internal/handler"
	"net/http"
	"time"
)

func main() {
	serverHandler := handler.NewHandler()

	if err := run("5000", serverHandler.InitRoutes()); err != nil {
		log.Fatal("error")
	}
}

func run(port string, handler http.Handler) error {
	server := http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 2,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return server.ListenAndServe()
}