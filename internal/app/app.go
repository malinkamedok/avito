package app

import (
	"avito/config"
	v1 "avito/internal/controller/http/v1"
	"avito/internal/usecase"
	"avito/internal/usecase/repo"
	"avito/pkg/httpserver"
	"avito/pkg/postgres"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {
	pg, err := postgres.New(cfg)
	if err != nil {
		log.Fatal("Cannot connnect to Postgre")
	}

	us := usecase.NewUserUseCase(repo.NewUserRepo(pg))

	handler := gin.New()
	
	v1.NewRouter(handler,
		us)

	serv := httpserver.New(handler, httpserver.Port(cfg.AppPort))
	interruption := make(chan os.Signal, 1)
	signal.Notify(interruption, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interruption:
		log.Printf("signal: " + s.String())
	case err = <-serv.Notify():
		log.Printf("Notify from http server")
	}

	err = serv.Shutdown()
	if err != nil {
		log.Printf("Http server shutdown")
	}

}
