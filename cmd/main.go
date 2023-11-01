package main

import (
	"messenger-go/internal/handler"
	"messenger-go/internal/logger"
	"messenger-go/internal/repository"
	"messenger-go/internal/repository/postgres"
	"messenger-go/internal/service"
	"net/http"
	"time"
)

func main() {
	messageRepo := postgres.NeMessagePostgres()

	repo := repository.NewRepository(messageRepo)
	services := service.NewService(repo)

	serverHandler := handler.NewHandler(services)

	if err := run("5000", serverHandler.InitRoutes()); err != nil {
		logger.Fatal("error")
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
